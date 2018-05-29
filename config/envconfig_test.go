package config

import (
	"fmt"
	"testing"
)

func TestReadConfigFromJSON(t *testing.T) {
	configFilePath := "../config.json"

	err := readConfigFromJSON(configFilePath)

	if err != nil {
		t.Errorf("Reading configuration failed: %s\n", err)
	}
}

func TestReadConfigFromJSONWrongPath(t *testing.T) {
	configFilePath := ""

	err := readConfigFromJSON(configFilePath)

	if err == nil {
		t.Errorf("Something's going wrong: %s\n", err)
	}
}

func TestReadConfigFromENV(t *testing.T) {
	err := readConfigFromENV()

	if err != nil {
		t.Error("Configuration is missing")
	}
}

func TestLoadReadingFromENV(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Error reading from ENV: %s", r)
		}

	}()

	ConfigFilePath := ""
	fmt.Printf("Reading from ENV with epty path (%s)", ConfigFilePath)

	Load()
}

func TestLoadReadingFromJSON(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Error reading from JSON: %s", r)
		}

	}()

	ConfigFilePath := "../config.json"
	fmt.Printf("Reading from JSON (%s)", ConfigFilePath)

	Load()
}
