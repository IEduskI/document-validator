package utils

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Configuration interface {
	GetServiceDocumentsTypes() []string
}

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) GetServiceDocumentsTypes() []string {
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	var model Model
	if err := yaml.NewDecoder(file).Decode(&model); err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}

	types := strings.Split(model.Service.Documents.Types, ",")

	return types
}
