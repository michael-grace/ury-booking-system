package config

import (
	"encoding/json"
	"io/ioutil"
)

type configuration struct {
	LogFile  string `json:"logFile"`
	Port     int    `json:"port"`
	Database struct {
		Host   string `json:"db-host"`
		Port   int    `json:"db-port"`
		User   string `json:"db-user"`
		Pass   string `json:"db-pass"`
		DBname string `json:"db-dbname"`
	} `json:"database"`
}

// Config contains the parameters for the booking system
var Config configuration

// ConfigurationSetup reads the config json file, and puts it in the struct
// Errors need panicing becuase the log hasn't started yet
func ConfigurationSetup() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &Config)
	if err != nil {
		panic(err)
	}

	Database, err = NewDatabaseConnection()
	if err != nil {
		panic(err)
	}
}
