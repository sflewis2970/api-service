package config

import (
	"log"
	"os"
	"testing"
)

func setEnvVars() error {
	var envVarsErr error

	// Set system environment variables
	envVarsErr = os.Setenv(ENV, "")
	if envVarsErr != nil {
		log.Print("Error getting system env variable")
		return envVarsErr
	}

	envVarsErr = os.Setenv(HOST, "")
	if envVarsErr != nil {
		log.Print("Error getting system host variable")
		return envVarsErr
	}

	envVarsErr = os.Setenv(PORT, "8080")
	if envVarsErr != nil {
		log.Print("Error getting system port variable")
		return envVarsErr
	}

	// Set redis cache server environment variables
	envVarsErr = os.Setenv(REDIS_TLS_URL, "cache")
	if envVarsErr != nil {
		log.Print("Error getting redis server TLS URL env variable")
		return envVarsErr
	}

	envVarsErr = os.Setenv(REDIS_URL, "cache")
	if envVarsErr != nil {
		log.Print("Error getting redis server URL env variable")
		return envVarsErr
	}

	envVarsErr = os.Setenv(REDIS_PORT, "6379")
	if envVarsErr != nil {
		log.Print("Error getting redis server port env variable")
		return envVarsErr
	}

	return nil
}

func TestGetData(t *testing.T) {
	// Set environment variables
	var envVarsErr error
	envVarsErr = setEnvVars()
	if envVarsErr != nil {
		t.Errorf("setEnvVars(): returned expected error, got %v", envVarsErr.Error())
		return
	}

	// Create config object
	cfg := Get()
	action := UPDATE_CONFIG_DATA

	// Test cases
	testCases := []struct {
		args string
	}{
		{args: action},
		{args: ""},
	}

	for _, tt := range testCases {
		_, cfgDataErr := cfg.GetData(tt.args)

		if cfgDataErr != nil {
			t.Errorf("GetData(%v): error not expected, got %v", tt.args, cfgDataErr)
		}
	}
}

func BenchmarkGetData(b *testing.B) {
	cfg := Get()

	// benchmark
	for idx := 0; idx < b.N; idx++ {
		_, cfgDataErr := cfg.GetData()
		if cfgDataErr != nil {
			b.Errorf("GetData(): unexpected error occurred, %v", cfgDataErr.Error())
		}
	}
}
