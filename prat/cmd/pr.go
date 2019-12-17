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

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "get the given PR",
	Long: `recover the data for the specified PR.
Note: PR id's are unique to account/region so we don't need to specify repo`,
	Run: func(cmd *cobra.Command, args []string) {
		ccSession := core.GetSession("default", "us-west-2")
		prId, _ := cmd.Flags().GetString("pr-id")
		core.GetPullRequest(ccSession, prId)
	},
}

func init() {
	rootCmd.AddCommand(prCmd)
	prCmd.Flags().StringP("pr-id", "i", "", "This is the system generated PR id")
}
