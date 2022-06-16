package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const CONFIG_FILE_NAME = "./config/config.json"

type ConfigData struct {
	HostName      string `json:"hostname"`
	HostPort      string `json:"hostport"`
	DataStoreName string `json:"dsname"`
	DataStorePort string `json:"dsport"`
}

type config struct {
	cfgData *ConfigData
}

var cfg *config

func (c *config) readConfigFile() error {
	data, readErr := ioutil.ReadFile(CONFIG_FILE_NAME)
	if readErr != nil {
		return readErr
	}

	c.cfgData = new(ConfigData)
	unmarshalErr := json.Unmarshal(data, c.cfgData)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	return nil
}

func (c *config) GetConfigData() *ConfigData {
	return c.cfgData
}

func GetConfig() *config {
	if cfg == nil {
		log.Print("creating config object")
		cfg = new(config)

		readErr := cfg.readConfigFile()
		if readErr != nil {
			log.Print("Error reading config file: ", readErr)
		}
	}

	log.Print("returning config object")
	return cfg
}
