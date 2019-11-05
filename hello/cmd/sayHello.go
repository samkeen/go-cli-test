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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sayHello = &cobra.Command{
	Use:   "hello",
	Short: "I say hello",
	Long: `You can define name in an ENV var of HELLO_NAME, 
or a config file with a name: key --config =config.yml
or the -name (-n_ flag)`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if viper.GetString("name")!=""{
			name = viper.GetString("name")
		}
		if viper.GetString("HELLO_NAME")!=""{
			name = viper.GetString("HELLO_NAME")
		}
		fmt.Println("Hello " + name)
	},
}

func init() {
	rootCmd.AddCommand(sayHello)
	sayHello.Flags().StringP("name", "n", "Anonymous", "Set your name")
}
