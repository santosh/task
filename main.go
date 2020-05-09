package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/santosh/taskmanager/cmd"
	"github.com/santosh/taskmanager/db"
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
