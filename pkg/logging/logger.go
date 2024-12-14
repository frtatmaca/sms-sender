package logging

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z"
)

var clusterName = os.Getenv("CLUSTER_NAME")

func NewLoggerWithLevel(output string, level string) *zap.SugaredLogger {
	lvl := zap.AtomicLevel{}
	err := lvl.UnmarshalText([]byte((level)))
	if err != nil {
		log.Fatalf("couldn't create logger, err:%s", err)
	}
	cfg := zap.Config{
		Level:       lvl,
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "sourceLocation",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout(timeFormat),
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{output},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    map[string]interface{}{"cluster": clusterName},
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.Sugar()
}
