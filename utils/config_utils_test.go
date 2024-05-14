package utils

import (
	"testing"
)

const ConfigPath = "../configs/app_test_config.toml"

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig(ConfigPath)
	if err != nil {
		t.Fatalf("Error loading config: %s", err)
	}

	if config.Business.Name != "Test Business" {
		t.Errorf("Expected business name to be 'Test Business', got '%s'", config.Business.Name)
	}

	if config.Business.StartTime != "08.00" {
		t.Errorf("Expected start time to be '08.00', got '%s'", config.Business.StartTime)
	}

	if config.Business.EndTime != "21.00" {
		t.Errorf("Expected end time to be '21.00', got '%s'", config.Business.EndTime)
	}

	if config.Business.MinModuleTime != 60 {
		t.Errorf("Expected min module time to be 60, got %d", config.Business.MinModuleTime)
	}

	if len(config.Business.Offices) != 2 {
		t.Errorf("Expected 2 offices, got %d", len(config.Business.Offices))
	}

	if config.Business.Offices[0] != "OFF_01" {
		t.Errorf("Expected first office to be 'OFF_01', got '%s'", config.Business.Offices[0])
	}

	if config.Business.Offices[1] != "OFF_02" {
		t.Errorf("Expected second office to be 'OFF_02', got '%s'", config.Business.Offices[1])
	}
}
