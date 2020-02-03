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
	Short: "list open PRs for a given repo",
	Long: `You can define name in an ENV var of PRAT_NAME, 
or a config file with a name: key --config =config.yml
or the -name (-n_ flag)`,
	Run: func(cmd *cobra.Command, args []string) {
		repoName, _ := cmd.Flags().GetString("repo-name")
		awsRegion, _ := cmd.Flags().GetString("aws-region")
		awsProfile, _ := cmd.Flags().GetString("aws-profile")
		ccSession := core.GetSession(awsProfile, awsRegion)
		core.ListPullRequest(ccSession, repoName)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("repo-name", "r", "", "Repository name where the PR resides")
	listCmd.Flags().StringP("aws-profile", "p", "default", "the AWS profile to use")
	listCmd.Flags().StringP("aws-region", "e", "us-west-2", "the AWS profile to use")
}
