package config

import (
	"github.com/BurntSushi/toml"
)

type DB struct {
	User     string
	Password string
	Host     string
	Port     uint16
	DBName   string
}

type Server struct {
	Port uint16
}

type Config struct {
	DB     DB
	Server Server
}

var conf Config

func Parse() error {
	_, err := toml.DecodeFile("config.toml", &conf)
	return err
}

func Get() *Config {
	return &conf
}
