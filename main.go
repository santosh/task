package main

import (
	"fmt"
	"os"

	"github.com/santosh/taskmanager/cmd"
)

func main() {
	doOrQuit(cmd.RootCmd.Execute())
}

func doOrQuit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
