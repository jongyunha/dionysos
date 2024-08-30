package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// Config structure for the YAML file
type Config struct {
	Docker struct {
		Image        string `yaml:"image"`
		Tag          string `yaml:"tag"`
		Interval     int    `yaml:"interval"`
		IntervalUnit string `yaml:"interval_unit"`
		Concurrent   bool   `yaml:"concurrent"`
	} `yaml:"docker"`
}

// LoadConfig reads the YAML configuration file and returns a Config struct
func LoadConfig(filePath string) (*Config, error) {
	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML file: %v", err)
	}

	return &config, nil
}