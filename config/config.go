package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Authentication struct {
	APIKey string `yaml:"apikey"`
}
type Defaults struct {
	Region            string `yaml:"region"`
	Image             string `yaml:"image"`
	Size              string `yaml:"size"`
	Key               string `yaml:"key"`
	Backups           bool   `yaml:"backups"`
	IPv6              bool   `yaml:"ipv6"`
	PrivateNetworking bool   `yaml:"private_networking"`
}

type Config struct {
	Authentication Authentication `yaml:"authentication"`
	Defaults       Defaults       `yaml:"defaults"`
}

const (
	defaultDirectory  = ".godo-cli"
	defaultConfigName = "config.yaml"
)

func GetDefaultDirectory() string {
	return defaultDirectory
}
func GetDefaultConfigName() string {
	return defaultConfigName
}

func GetConfigDirectory() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", fmt.Errorf("Error Getenv $HOME not found")
	}

	return filepath.Join(home, defaultDirectory), nil
}

func GetConfigPath() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", fmt.Errorf("Error Getenv $HOME not found")
	}

	return filepath.Join(home, defaultDirectory, defaultConfigName), nil
}

func LoadConfig(path string) (*Config, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading %s: %s", path, err)
	}

	config := &Config{}

	err = yaml.Unmarshal(d, config)
	if err != nil {
		return nil, fmt.Errorf("Error yaml.Unmarshal %s: %s", path, err)
	}

	return config, nil
}

func SaveConfig(path string, config *Config) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("Error yaml.Marshal  %s", err)
	}

	err = ioutil.WriteFile(path, data, 0644)

	if err != nil {
		return fmt.Errorf("Error yaml.Marshal  %s", err)
	}

	return nil
}
