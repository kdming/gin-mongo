package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DB_HOST  string `yaml:"DB_HOST"`
	DB_USER  string `yaml:"DB_USER"`
	DB_PWD   string `yaml:"DB_PWD"`
	DB_NAME  string `yaml:"DB_NAME"`
	TOKENKEY string `yaml:"TOKENKEY"`
}

func GetConfig() Config {
	root, err := os.Getwd()
	filePath := root + "/config.yaml"

	buffer, err := ioutil.ReadFile(filePath)
	config := Config{}
	err = yaml.Unmarshal(buffer, &config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return config
}
