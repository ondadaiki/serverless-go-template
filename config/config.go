package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Postgres *DbConfig
}

type DbConfig struct {
	Host     string `json:host`
	User     string `json:user`
	DbName   string `json:dbname`
	Password string `json:password`
}

func NewConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	return &config, nil
}

func (config *DbConfig) ConnectionString() string {
	return fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable host=%s",
		config.User,
		config.DbName,
		config.Password,
		config.Host,
	)
}
