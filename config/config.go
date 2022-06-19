package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Directory struct {
	Format   string `yaml:"format"`
	Symbol   string `yaml:"symbol"`
	Disabled bool   `yaml:"disabled"`
}

type GitBranch struct {
	Format   string `yaml:"format"`
	Symbol   string `yaml:"symbol"`
	Disabled bool   `yaml:"disabled"`
}

type Config struct {
	Directory Directory `yaml:"directory"`
	GitBranch GitBranch `yaml:"git_branch"`
}

func NewConfig() (Config, error) {
	config := Config{}
	// Default configuration
	defaultConfig, err := DefaultConfig.ReadFile("config.yaml")

	if err != nil {
		return config, err
	}

	yaml.Unmarshal(defaultConfig, config)

	// User configuration
	userConfig, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		return config, err
	}

	yaml.Unmarshal(userConfig, config)

	return config, nil
}
