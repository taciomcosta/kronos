<p align="center">
  <img src="./docs/kronos.png">
  <br><br>
  Kronos is a cross-platform job scheduler that helps you manage, monitor and inspect cronjobs.
</p>

---

[![GitHub Actions](https://github.com/taciomcosta/kronos/workflows/main/badge.svg)](https://github.com/ktr0731/evans/actions)
[![codecov](https://codecov.io/gh/taciomcosta/kronos/branch/master/graph/badge.svg?token=bVwkHbOHOk)](https://codecov.io/gh/taciomcosta/kronos)
[![Go Report Card](https://goreportcard.com/badge/github.com/taciomcosta/kronos)](https://goreportcard.com/report/github.com/taciomcosta/kronos)
[![Latest Release](https://img.shields.io/github/v/release/taciomcosta/kronos?include_prereleases)](https://github.com/taciomcosta/kronos/releases/latest)

**This project is being developed using [RDD](https://tom.preston-werner.com/2010/08/23/readme-driven-development.html) and is not production-ready yet!**
<br>Check the [architecture](docs/ARCHITECTURE.md) of the project and feel free to contribute. :)

## Table of Contents
- [Installation](#Installation)
  - [Linux](#Linux)
  - [macOS](#macOS)
  - [Windows](#Windows)
- [Commands](#Commands)
- [Examples](#Examples)
  - [Jobs](#Jobs)
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
brew tap taciomcosta/kronos
brew install kronos
sudo brew services start kronos
```

### Windows
```
choco install kronos
```

By default, Kronos daemon uses `:8080` as its host and creates a db named
`kronos.db`. <br>
You can create a file at `<home-directory>/.kronosd/config.json`
to override this behavior.
```json
{
    "host": ":8080",
    "db": "kronos.db"
}
```


## Commands
```
> kronos
kronos is a cross-platform job scheduler that helps you manage, monitor and inspect cronjobs.

Usage:
  kronos [command]

Available Commands:
  create         Creates a new job/channel
  delete         Deletes a job/channel
  list           Lists all jobs/channels
  describe       Shows detailed information about a job/channel
  history        Shows execution history of a job
  attach         Attaches local stdin, stdout, stderr to a job
  logs           Prints logs for a specific job execution
  enable         Enables a job execution
  disable        Disables a job execution
  assign         Assigns a channel to a job

Flags:
  -h, --help   help for kronos

Use "kronos <command> --help" to learn more about a specific command.
```

## Examples

Creating a new job with sugar expressions: 
```
> kronos create job --name myjob --cmd ./my-job.sh --tick "every day"
```
Alternatively, we can use regular cron expressions:
```
> kronos create job --name myjob --cmd ./my-job.sh --tick "0 0 */1 * *"
```

Listing jobs:
```
> kronos list jobs
NAME            COMMAND             TICK               LAST EXECUTION             STATUS
myjob           ./my-job.sh         Every day          2021-01-01 00:00:00        Enabled 
myfiles         ls                  0 0 */1 * *        2021-01-01 00:00:00        Disabled
backup-db       ./backup            Every day          2021-01-01 00:00:00        Enabled
```

Describing a job:
```
> kronos describe job myjob
Name: myjob
Command: ./my-job.sh
Tick: Every day
Last Execution: 2021-01-01 00:00:00
Status: Enabled

Executions: 4 
 - 3 Succeeded
 - 1 Failed

Resources:
 - Average CPU: 50%
 - Average Memory: 300MB
 - Average Networking: 100MB (IN) / 50KB (OUT)
 
Assigned Channels:
  - my-slack
  - my-email
```

Attaching to a job:
```
> kronos attach myjob
Attached to myjob (Press CTRL+C to exit)

$ Doing the thing...
$ Job finished
```

Alternatively, we can just log the last execution instead of attaching to the job:
```
> kronos logs  myjob
$ Doing the thing...
$ Job finished
> 
```

Showing execution history for all jobs:
```
> kronos history
NAME            EXECUTION                  STATUS
myjob           2021-01-01 00:00:00        Succeeded
myfiles         2021-01-02 00:00:00        Failed
backup-db       2021-01-03 00:00:00        Succeeded
```

Last 3 executions for a specific job:
```
> kronos history myjob --last 3
NAME            EXECUTION                  STATUS
myjob           2021-01-01 00:00:00        Failed
myjob           2021-01-02 00:00:00        Succeeded
myjob           2021-01-03 00:00:00        Succeeded
```


Deleting a job:
```
> kronos delete job myjob
myjob deleted.
```

