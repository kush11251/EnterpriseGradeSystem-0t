package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config represents the application configuration
//
// @Description: Config is a struct that holds the application configuration.
// @Fields:
//   - Database: string
//   - Server: ServerConfig
//
type Config struct {
	Database string `json:"database"`
	Server   ServerConfig `json:"server"`
}

// ServerConfig represents the server configuration
//
// @Description: ServerConfig is a struct that holds the server configuration.
// @Fields:
//   - Port: int
//   - Host: string
//
type ServerConfig struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

func Init() error {
	configFile, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer configFile.Close()

	var config Config
	configBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(configBytes, &config); err != nil {
		return err
	}

	return nil
}