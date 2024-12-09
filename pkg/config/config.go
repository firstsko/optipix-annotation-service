package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port     int    `yaml:"port"`
	LogLevel string `yaml:"log_level"`
}

type WebSocketConfig struct {
	MaxMessageSize  int `yaml:"max_message_size"`
	ReadBufferSize  int `yaml:"read_buffer_size"`
	WriteBufferSize int `yaml:"write_buffer_size"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Config struct {
	Server    ServerConfig    `yaml:"server"`
	WebSocket WebSocketConfig `yaml:"websocket"`
	Database  DatabaseConfig  `yaml:"database"`
}

var AppConfig *Config = &Config{}

func LoadConfig(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(AppConfig); err != nil {
		log.Fatalf("failed to decode config file: %v", err)
	}

	log.Printf("Config loaded")
}
