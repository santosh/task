# cotu

`cotu`, an abbreviation for *Center of the Universe* is an experimental CLI tool to connect all the services a person can use daily.

## Features

- Connects to different kinds of API backends, including
  - REST
  - GraphQL

## Usage

```text
$ cotu
Center of the Universe.

A command line tool to manage your daily task.

Usage:
  cotu [command]

Available Commands:
  help        Help about any command
  task        A local db based task manager
  version     Prints the version number of cotu

Flags:
      --config string   config file (default is $HOME/.task.yaml)
  -h, --help            help for cotu
  -t, --toggle          Help message for toggle

Use "cotu [command] --help" for more information about a command.
```

## Development

As I am a regular user of these services, I will be implementing these third-party services in this CLI tool.

- Twitter
- Reddit
- Github
- Slack
- Gmail?
- Google Calendar?
- LinkedIn

### Design Goals

- For every subcommand, there will be an `auth` subcommand.
- All other subcommands for that subcommand will depend on auth middleware.
