# cotu

`cotu`, an abbreviation for *Center of the Universe* is an experimental CLI tool to connect all the services a person can use daily.

## Installation

    go get -u github.com/santosh/cotu

## Highlights

<!-- This section only lists non-technical writeup -->

- Connects to different ever changing genres of backend API, including
  - REST
  - GraphQL
  - Twitter's **Streaming API**
- System is extensible to

<!-- More technical writeup to be written on the blog

Is object-oriented.
Follows twelve-factor methodology (don't wait until all are achieved)
Follows different authentication mechanism, OAuth, JWT
Follows different API mechanism

-->

<!-- Post to HackeNews, r/Golang -->

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
