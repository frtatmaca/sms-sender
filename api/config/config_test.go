package config_test

import (
	"os"
	"testing"

	"github.com/frtatmaca/sms-sender/api/config"
)

func TestConfig_DevYamlConfig(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("APP_ENV", "mock")

	tests := map[string]struct {
		expected config.AppConfig
	}{
		"should map config": {
			expected: config.AppConfig{
				Api: config.ApiConfig{
					AppId:   "sms-sender",
					AppName: "sms-sender-api",
				},
			},
		},
	}

	for name, tc := range tests {
		result := config.NewConfiguration()
		if tc.expected.Api.AppName != result.Api.AppName {
			t.Fatalf("expected error: %v", name)
		}
	}
}
