package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	APPPORT        string `yaml:"APPPORT"`
	SaveRequestLog bool   `yaml:"SaveRequestLog"`
	MongoUrl       string `yaml:"MongoUrl"`
	DB_NAME        string `yaml:"DB_NAME"`
	Token_KEY      string `yaml:"Token_KEY"`
	HttpUrl        string `yaml:"HttpUrl"`
}

func GetConfig() Config {
	// 设置文件路径
	root, err := os.Getwd()
	filePath := root + "/config.yaml"

	// 读取并解析文件
	buffer, err := ioutil.ReadFile(filePath)
	config := Config{}
	err = yaml.Unmarshal(buffer, &config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return config
}
