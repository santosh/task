# taskmanager

Task Manager is an experimental CLI task management app written in Golang and BoltDB. Task Manager is not to be confused with the system resource manager.

## Usage

```text
Task is a CLI task manager

Usage:
  task [command]

Available Commands:
  add         Adds a task to your task list.
  do          Marks a task as complete.
  help        Help about any command
  list        Lists all of your tasks.

Flags:
      --config string   config file (default is $HOME/.task.yaml)
  -h, --help            help for task
  -t, --toggle          Help message for toggle

Use "task [command] --help" for more information about a command.
```

