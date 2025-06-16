package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type AdminConfig struct {
	Env string `yaml:"env"`
	HTTPServer `yaml:"hhtp_server"`
	Auth	   `yaml:"auth"`
}

type HTTPServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Auth struct {
	Secret string `yaml:"secret"`
}

func MustLoad() *AdminConfig {
	// Processing env config variable and file
	configPath := os.Getenv("ADMIN_CONFIG_PATH")
	if configPath == ""{
		log.Fatalf("ADMIN_CONFIG_PATH was not found\n")
	}

	if _, err := os.Stat(configPath); err != nil{
		log.Fatalf("failed to find config file: %v\n", err)
	}

	// YAML to struct object
	var cfg AdminConfig
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil{
		log.Fatalf("failed to read config file: %v", err)
	}

	return &cfg
}