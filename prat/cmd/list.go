/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"cli-go-test/prat/core"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list from Prat",
	Long: `You can define name in an ENV var of PRAT_NAME, 
or a config file with a name: key --config =config.yml
or the -name (-n_ flag)`,
	Run: func(cmd *cobra.Command, args []string) {
		core.ListRepos("default", "us-west-2")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	//listCmd.Flags().StringP("name", "n", "Anonymous", "Say your name")
}
