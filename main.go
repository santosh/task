package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/santosh/task/cmd"
	"github.com/santosh/task/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	doOrQuit(db.Init(dbPath))
	defer db.Close()
	doOrQuit(cmd.RootCmd.Execute())
}

func doOrQuit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
