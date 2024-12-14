package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type IAppConfig interface {
	ReadAppConfig()
}

type AppConfig struct {
	v              *viper.Viper
	Api            ApiConfig   `required:"true" mapstructure:"api"`
	LogLevel       string      `required:"true" yaml:"loglevel"`
	Port           int         `required:"true" yaml:"port"`
	CronExpression string      `required:"true" yaml:"cronExpression"`
	Redis          RedisConfig `required:"true" mapstructure:"redis"`
}

type ApiConfig struct {
	AppEnv  string `required:"true" mapstructure:"appEnv"`
	AppId   string `required:"true" mapstructure:"appId"`
	AppName string `required:"true" mapstructure:"name"`
}

type RedisConfig struct {
	Addr string `required:"true" mapstructure:"addr"`
}

func (c *AppConfig) ReadAppConfig() {
	v := viper.New()

	env := strings.ToLower(os.Getenv("APP_ENV"))

	if env == "" {
		env = "local"
	}

	v.SetTypeByDefaultValue(true)
	v.SetConfigName(env)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("envs")
	v.AddConfigPath(filepath.Dir(env))

	c.v = v

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(c); err != nil {
		panic(err)
	}
}

func NewConfiguration() *AppConfig {
	applicationConfig := &AppConfig{}
	applicationConfig.ReadAppConfig()
	applicationConfig.v.WatchConfig()
	applicationConfig.v.OnConfigChange(func(in fsnotify.Event) {
		applicationConfig.ReadAppConfig()
	})

	return applicationConfig
}
