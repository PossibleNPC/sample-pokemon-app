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

// NewConfig is meant to return a config representing configuration of our app. It requires a filepath to a yaml file
// that conforms to the struct above. Config value priority from low to high is: default app yaml = yaml file flag ->
// os.ENV
func NewConfig(filepath string) *Config {
	cfg := &Config{}

	cfg.getConfFile(filepath)
	cfg.getConfEnv()
	cfg.dsn()

	return cfg
}

func (cfg *Config) getConfFile(filepath string) {
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

func (cfg *Config) getConfEnv() {
	err := envconfig.Process("", cfg)
	if err != nil {
		confError(err)
	}
}

func (cfg *Config) dsn() {
	cfg.Database.Dsn = fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port)
}

// TODO: This might need to be gotten rid of
func confError(err error) {
	fmt.Println("received error: ", err)
	os.Exit(1)
}