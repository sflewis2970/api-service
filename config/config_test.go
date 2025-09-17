package config

import (
	"fmt"
	"log"
	"testing"
)

func TestGetData(t *testing.T) {
	// Set environment variables
	var envVarsErr error
	envVarsErr = SetEnvVars()
	if envVarsErr != nil {
		t.Errorf("setEnvVars(): returned an error, got %v", envVarsErr.Error())
		return
	}

	// Create config object
	cfg := CreateNewConfigObject()
	if cfg == nil {
		t.Errorf("error could not create cfg object.\n")

	} else {
		_, cfgDataErr := GetConfigData(t, cfg)
		if cfgDataErr != nil {
			t.Errorf("error getting cfg data, with error: %s\n", cfgDataErr.Error())
		}
	}
}

func GetConfigData(t *testing.T, cfg *Config) (*CfgData, error) {
	cfgData, cfgDataErr := cfg.GetData(REFRESH_CONFIG_DATA)

	if cfgData == nil {
		t.Errorf("nil object returned getting config data...")
	}

	if cfgDataErr != nil {
		t.Errorf("Error getting config data with error: %s", cfgDataErr.Error())
		return nil, cfgDataErr
	} else {
		log.Printf("getting config data successful: %v", cfgData)
		return cfgData, nil
	}
}

func TestGetDataWithTestCases(t *testing.T) {
	// Create config object
	cfg := CreateNewConfigObject()

	// Set environment variables
	var envVarsErr error
	envVarsErr = SetEnvVars()
	if envVarsErr != nil {
		t.Errorf("setEnvVars(): returned expected error, got %v", envVarsErr.Error())
		return
	}

	action := REFRESH_CONFIG_DATA

	// Test cases
	testCases := []struct {
		args string
	}{
		{args: action},
		{args: ""},
	}

	for _, tt := range testCases {
		cfgData, cfgDataErr := cfg.GetData(tt.args)

		if cfgDataErr != nil {
			t.Errorf("GetData(%v): error not expected, got %v", tt.args, cfgDataErr)
		} else {
			fmt.Printf("config data: %v", cfgData)
		}
	}
}

func TestNewConfig(t *testing.T) {
	cfg := CreateNewConfigObject()

	if cfg == nil {
		t.Errorf("Could not create Config object!")
	} else {
		fmt.Println("Config object created!")
	}
}

func CreateNewConfigObject() *Config {
	cfg := New()

	return cfg
}
func BenchmarkGetData(b *testing.B) {
	cfg := CreateNewConfigObject()

	// benchmark
	for idx := 0; idx < b.N; idx++ {
		_, cfgDataErr := cfg.GetData()
		if cfgDataErr != nil {
			b.Errorf("GetData(): unexpected error occurred, %v", cfgDataErr.Error())
		}
	}
}
