package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/sntshk/task/db"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			log.Fatal("Something went wrong: ", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			log.Fatal("You have no tasks to complete!")
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
