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

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create PRs",
	Long: `command to create a pull request`,
	Run: func(cmd *cobra.Command, args []string) {
		ccSession := core.GetSession("default", "us-west-2")
		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("desc")
		srcBranch, _ := cmd.Flags().GetString("src-branch")
		destBranch, _ := cmd.Flags().GetString("dest-branch")
		repoName, _ := cmd.Flags().GetString("repo-name")
		core.CreatePullRequest(ccSession, title, description, srcBranch, destBranch, repoName)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("title", "t", "", "Tile for the PR")
	createCmd.Flags().StringP("desc", "d", "", "Description for the PR")
	createCmd.Flags().StringP("src-branch", "s", "", "Source branch for the PR")
	createCmd.Flags().StringP("dest-branch", "e", "", "Destination branch for the PR")
	createCmd.Flags().StringP("repo-name", "r", "", "Repository name where the PR resides")
}
