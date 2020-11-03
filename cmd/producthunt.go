/*
Copyright © 2020 Santosh Kumar <sntshkmr60@gmail.com>

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
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func init() {
	home, _ := homedir.Dir()
	dbPath = filepath.Join(home, "tasks.db")
}

// productHuntCmd represents the prodhunt command
var productHuntCmd = &cobra.Command{
	Use:   "prodhunt",
	Short: "show trending products from producthunt.com",
}

// listCmd represents the list command
var topCmd = &cobra.Command{
	Use:   "top",
	Short: "list top 3 products of this month",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Showing top products...")
	},
}

func init() {
	productHuntCmd.AddCommand(topCmd)
	RootCmd.AddCommand(productHuntCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
