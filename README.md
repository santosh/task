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
  github      interface with github
  help        Help about any command
  reddit      interface with reddit
  s3          interface with cotu s3 bucket
  task        a local db based task manager
  twitter     interface with twitter
  version     prints the version number of cotu

Flags:
  -h, --help   help for cotu
```

## Development

Dummy commands have been added for:

- **task** - a BoltDB based task manager
    - *add*: add new task
    - *do*: mark task as done
    - *list*: list all pending task
- **s3** - an interface to cotu s3 bucket
    - *upload*: upload file to cotu bucket
    - *list*: list files from cotu bucket
- **twitter** - commands for twitter
- **reddit** - commands for reddit
- **github** - commands for github
- **slack** - commands for slack
- **calendar** - commands for google calendar

### Design Goals

- For every subcommand, there will be an `auth` subcommand.
- All other subcommands for that subcommand will depend on auth middleware.
