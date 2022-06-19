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

	"github.com/lab42/rocketship/shell"
	"github.com/spf13/cobra"
)

func zsh(cmd *cobra.Command, args []string) {
	shellScriptBytes, err := shell.ShellFiles.ReadFile("zsh.shell")
	cobra.CheckErr(err)
	fmt.Println(string(shellScriptBytes))
}

// zshCmd represents the zsh command
var zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Print shell functions to init Rocketship for Zsh.",
	Run:   zsh,
}

func init() {
	initCmd.AddCommand(zshCmd)
}
