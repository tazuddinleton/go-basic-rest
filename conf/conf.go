package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	dlog "log"
)

const configFilePath string = "config.json"

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

// Logging returns logging configuration object
func (c config) Logging() logging {
	return c.logging
}

// Db returns database configuration object
func (c config) Db() db {
	return c.db
}
func (db db) ConnString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.User, db.Password, db.Host, db.Port, db.Database)
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
