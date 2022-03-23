package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Database struct {
		User string `yaml:"user" envconfig:"DB_USER"`
		Password string `yaml:"password" envconfig:"DB_PASSWORD"`
		Database string `yaml:"database" envconfig:"DB"`
		Host string `yaml:"host" envconfig:"DB_HOST"`
		Port int `yaml:"port" envconfig:"DB_PORT"`
	} `yaml:"postgres"`
	Api struct {
		Port int `yaml:"port"`
	}`yaml:"api"`
}

func GetConfFile(cfg *Config) {
	ymlConf, err := os.Open("./config.yml")
	if err != nil {
		confError(err)
	}
	defer ymlConf.Close()

	decoder := yaml.NewDecoder(ymlConf)
	err = decoder.Decode(&cfg)
	if err != nil {
		confError(err)
	}
}

func GetConfEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		confError(err)
	}
}

func confError(err error) {
	fmt.Println("received error: ", err)
	os.Exit(1)
}