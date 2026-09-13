package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Redis  RedisConfig  `yaml:"redis"`
}

type ServerConfig struct {
	Host string `yaml:"host" default:"http://localhost"`
	Port uint8  `yaml:"port" default:"8080"`
}

type RedisConfig struct {
	Addr       string `yaml:"addr" default:"localhost:6379"`
	Password   string `yaml:"password" default:""`
	DB         int    `yaml:"db" default:"0"`
	TTLSeconds int    `yaml:"ttl_seconds" default:"300"`
}

func LoadConfig(path string) (*Config, error) {
	const operation = "wx-api.internal.config.server.LoadConfig"

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("%s: fatal error: opening config file: %s", operation, err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalf("%s: fatal error: closing config file: %s", operation, err)
		}
	}()

	var config Config

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		log.Fatalf("%s: fatal error: decoding config file: %s", operation, err)
	}

	return &config, nil
}
