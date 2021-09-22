package conf

import (
	"encoding/json"
	"io/ioutil"
	dlog "log"
)

const configFilePath string = "config.json"

type Configuration interface {
	Logging() logging
	Db() db
}

type db struct {
}

func (c config) Logging() logging {
	return c.logging
}

func (c config) Db() db {
	return c.db
}

type config struct {
	Name    string
	Version string
	logging logging
	db      db
}

type configData struct {
	Name       string  `json:"name"`
	Version    string  `json:"version"`
	LoggingDef logging `json:"logging"`
	DbDef      db      `json:"db"`
}

type logging struct {
	Level          string `json:"level"`
	File           string `json:"file"`
	ConsoleEnabled bool   `json:"consoleEnabled"`
}

// New returns new Config object
func New() *config {
	data := readData()
	return &config{
		Name:    data.Name,
		Version: data.Version,
		logging: data.LoggingDef,
		db:      data.DbDef,
	}
}

func readData() configData {
	bytes, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		dlog.Fatal("Failed to read config.json")
	}
	var c configData
	err = json.Unmarshal(bytes, &c)

	if err != nil {
		dlog.Fatal("Failed to parse json")
	}
	return c
}

func (c *configData) String() string {
	return c.Name + ", Version: " + c.Version
}
