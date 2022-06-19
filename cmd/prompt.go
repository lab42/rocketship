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
	"sync"

	"github.com/lab42/rocketship/config"
	"github.com/lab42/rocketship/formatter"
	"github.com/lab42/rocketship/modules"
	"github.com/spf13/cobra"
)

func prompt(cmd *cobra.Command, args []string) {
	continuation, _ := cmd.Flags().GetBool("continuation")

	if continuation {
		config, err := config.NewConfig()
		cobra.CheckErr(err)

		result, err := modules.Modules["continuation"].Exec(&config)
		cobra.CheckErr(err)

		fmt.Print(result)
	}

	if !continuation {
		var data = make(map[string]string)

		config, err := config.NewConfig()
		cobra.CheckErr(err)

		var wg sync.WaitGroup
		var mutex = &sync.Mutex{}

		for _, module := range modules.Modules {
			wg.Add(1)
			go func(module modules.Module, data map[string]string) {
				defer wg.Done()

				result, err := module.Exec(&config)
				cobra.CheckErr(err)

				mutex.Lock()
				data[module.Name] = result
				mutex.Unlock()

			}(*module, data)
		}

		wg.Wait()

		data["line_break"] = "\n"

		formatter := formatter.NewFormatter(
			"cli_format",
			config.Format,
			data,
		)

		result, err := formatter.Render()
		cobra.CheckErr(err)

		fmt.Print(result)
	}
}

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Render the prompt",
	Run:   prompt,
}

func init() {
	rootCmd.AddCommand(promptCmd)
	promptCmd.Flags().BoolP("continuation", "t", false, "Help message for toggle")
}
