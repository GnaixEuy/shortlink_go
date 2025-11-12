package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	MySQL struct {
		DSN          string `yaml:"dsn"`
		MaxOpenConns int    `yaml:"max_open_conns"`
		MaxIdleConns int    `yaml:"max_idle_conns"`
	} `yaml:"mysql"`
	Redis struct {
		Addr      string `yaml:"addr"`
		Password  string `yaml:"password"`
		DB        int    `yaml:"db"`
		KeyPrefix string `yaml:"key_prefix"`
		TTL       int    `yaml:"ttl_seconds"`
	} `yaml:"redis"`
}

var C Config

func Load(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, &C)
}
