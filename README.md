![kronos](docs/kronos.png)

Kronos is a cross-platform job scheduler that helps you manage, monitor and inspect jobs.

[_add Gif showing usage_]

## Table of Contents
- [Installation](#Installation)
  - [Linux](#Linux)
  - [macOS](#macOS)
  - [Windows](#Windows)
- [Commands](#Commands)
- [Examples](#Examples)
  - [Jobs](#Jobs)
    - [Managing](#Managing)
    - [Inspecting](#Inspecting)
  - [Channels](#Channels)
    - [Slack](#Slack)
    - [E-mail](#Email)


## Installation
### Linux

**Debian/Ubuntu**
```
sudo apt-get install kronos
```

**Arch**
```
sudo pacman -S kronos
```

### macOs
```
$ brew tap taciomcosta/kronos
$ brew install kronos
```

### Windows
```
choco install kronos
```

## Commands
```
> kronos
kronos is a cross-platform job scheduler that helps you manage, monitor and inspect jobs.

Commands:
  create         Creates a new job/channel
  delete         Deletes a job/channel
  list           Lists all jobs/channels
  describe       Shows detailed information about a job/channel
  history        Shows execution history of a job
  attach         Attachs local stdin, stdout, stderr to a job
  logs           Prints logs for a specific job execution
  enable         Enables a job execution
  disable        Disables a job execution
  assign         Assigns a channel to a job

Use "kronos <command> --help" to know more about a specific command.
```



