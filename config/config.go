package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
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

type Golang struct {
	Format   string `yaml:"format"`
	Symbol   string `yaml:"symbol"`
	Disabled bool   `yaml:"disabled"`
}

type Time struct {
	Format   string `yaml:"format"`
	Symbol   string `yaml:"symbol"`
	Disabled bool   `yaml:"disabled"`
}

type Character struct {
	Format   string `yaml:"format"`
	Disabled bool   `yaml:"disabled"`
}

type Continuation struct {
	Format   string `yaml:"format"`
	Disabled bool   `yaml:"disabled"`
}

type Config struct {
	Format       string       `yaml:"format"`
	Directory    Directory    `yaml:"directory"`
	GitBranch    GitBranch    `yaml:"git_branch"`
	Golang       Golang       `yaml:"golang"`
	Time         Time         `yaml:"time"`
	Character    Character    `yaml:"character"`
	Continuation Continuation `yaml:"continuation"`
}

func NewConfig() (Config, error) {
	config := Config{}
	// Default configuration
	defaultConfig, err := DefaultConfig.ReadFile("config.yaml")

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(defaultConfig, &config)

	if err != nil {
		return config, err
	}

	// User configuration
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	configFile := fmt.Sprintf("%s/.config/rocketship.yaml", homeDir)

	userConfig, _ := ioutil.ReadFile(configFile)
	_ = yaml.Unmarshal(userConfig, &config)

	return config, nil
}
