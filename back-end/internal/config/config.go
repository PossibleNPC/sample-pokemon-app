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
		Port string `yaml:"port" envconfig:"DB_PORT"`
		Dsn string
	} `yaml:"postgres"`
	Api struct {
		Port string `yaml:"port"`
	}`yaml:"api"`
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) GetConfFile(filepath string) {
	ymlConf, err := os.Open(filepath)
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

func (cfg *Config) GetConfEnv() {
	err := envconfig.Process("", cfg)
	if err != nil {
		confError(err)
	}
}

func (cfg *Config) PsqlConnString() {
	cfg.Database.Dsn = fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port)
}

// TODO: This might need to be gotten rid of
func confError(err error) {
	fmt.Println("received error: ", err)
	os.Exit(1)
}