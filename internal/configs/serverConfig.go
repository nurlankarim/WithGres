package configs

import (
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
)

type Config struct {
	ServerAddress string `yaml:"server_address"`
	DatabaseURL   string `yaml:"db_url"`
}

func NewConfig() *Config {
	var cfg *Config
	bytes, err := ioutil.ReadFile("configs/application.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
