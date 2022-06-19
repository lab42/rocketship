/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/lab42/rocketship/config"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func configure(cmd *cobra.Command, args []string) {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	configFile := fmt.Sprintf("%s/.config/rocketship.yaml", homeDir)
	_, err = os.Stat(configFile)

	if os.IsExist(err) {
		defaultConfig, _ := config.DefaultConfig.ReadFile("config.yaml")
		os.WriteFile(configFile, defaultConfig, 0644)
	} else {
		prompt := promptui.Select{
			Label: "Configuration exists. Would you like to rewrite?",
			Items: []string{"yes", "no"},
		}

		_, result, _ := prompt.Run()

		if result == "yes" {
			defaultConfig, _ := config.DefaultConfig.ReadFile("config.yaml")
			os.WriteFile(configFile, defaultConfig, 0644)
			os.Exit(0)
		}

		if result == "no" {
			os.Exit(0)
		}
	}
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Generate and write default config file.",
	Run:   configure,
}

func init() {
	rootCmd.AddCommand(configCmd)
}
