package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/kelseyhightower/envconfig"
)

const (
	defaultConfigFilePath = "config.json"
)

var (
	// Config is a package variable, which is populated during loading and shared to whole application
	Config Configuration
	// ConfigFilePath defines a path to JSON-config file
	ConfigFilePath = defaultConfigFilePath
)

// Configuration options
type Configuration struct {
	HTTPListenURL    string `json:"HTTPListenURL"       envconfig:"BOW_API_HTTP_LISTEN_URL"     default:":8080"`
	HTTPReadTimeout  int    `json:"HTTPReadTimeout"     envconfig:"BOW_API_HTTP_READ_TIMEOUT"   default:"15"`
	HTTPWriteTimeout int    `json:"HTTPWriteTimeout"    envconfig:"BOW_API_HTTP_WRITE_TIMEOUT"  default:"15"`
	HTTPIdleTimeout  int    `json:"HTTPIdleTimeout"     envconfig:"BOW_API_HTTP_IDLE_TIMEOUT"   default:"15"`
	HTTPGraceTimeout int    `json:"HTTPGraceTimeout"    envconfig:"BOW_API_HTTP_GRACE_TIMEOUT"  default:"60"`
}

// Load reads and loads configuration to Config variable
func Load() {
	var err error

	confLen := len(ConfigFilePath)
	if confLen != 0 {
		err = readConfigFromJSON(ConfigFilePath)
	}
	if confLen == 0 || err != nil {
		err = readConfigFromENV()
	}
	if err != nil {
		panic(`Configuration not found. Please specify configuration`)
	}
}

// readConfigFromJSON reads config data from JSON-file
func readConfigFromJSON(configFilePath string) error {
	log.Printf("Looking for JSON config file (%s)", configFilePath)

	contents, err := ioutil.ReadFile(configFilePath)
	if err == nil {
		reader := bytes.NewBuffer(contents)
		err = json.NewDecoder(reader).Decode(&Config)
	}
	if err != nil {
		log.Printf("Reading configuration from JSON (%s) failed: %s\n", configFilePath, err)
	} else {
		log.Printf("Configuration has been read from JSON (%s) successfully\n", configFilePath)
	}

	return err
}

// readConfigFromENV reads data from environment variables
func readConfigFromENV() error {
	log.Println("Looking for ENV configuration")
	return envconfig.Process("BOW_API", &Config)
}
