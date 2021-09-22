package conf

import (
	"encoding/json"
	"io/ioutil"
	dlog "log"
)

const configFilePath string = "config.json"

type Config struct {
	Name       string
	Version    string
	ConfigData ConfigData
}

func (c Config) Logging() Logging {
	return c.ConfigData.Logging
}

type ConfigData struct {
	Name    string  `json:"name"`
	Version string  `json:"version"`
	Logging Logging `json:"logging"`
}

type Logging struct {
	Level          string `json:"level"`
	File           string `json:"file"`
	ConsoleEnabled bool   `json:"consoleEnabled"`
}

// New returns new Config object
func New() *Config {
	data := readData()
	return &Config{
		data.Name, data.Version, data,
	}
}
func readData() ConfigData {
	bytes, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		dlog.Fatal("Failed to read config.json")
	}
	var c ConfigData
	err = json.Unmarshal(bytes, &c)

	if err != nil {
		dlog.Fatal("Failed to parse json")
	}
	return c
}

func (c *Config) String() string {
	return c.Name + ", Version: " + c.Version
}
