package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gitlab.com/sntshk/task/cmd"
	"gitlab.com/sntshk/task/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	doOrQuit(db.Init(dbPath))
	doOrQuit(cmd.RootCmd.Execute())
}

func doOrQuit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
