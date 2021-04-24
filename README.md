<p align="center">
  <img src="./docs/kronos.png">
  <br><br>
  Kronos is a cross-platform job scheduler that helps you manage, monitor and inspect cronjobs.
</p>

---

[![GitHub Actions](https://github.com/taciomcosta/kronos/workflows/main/badge.svg)](https://github.com/taciomcosta/kronos/actions)
[![codecov](https://codecov.io/gh/taciomcosta/kronos/branch/master/graph/badge.svg?token=bVwkHbOHOk)](https://codecov.io/gh/taciomcosta/kronos)
[![Go Report Card](https://goreportcard.com/badge/github.com/taciomcosta/kronos)](https://goreportcard.com/report/github.com/taciomcosta/kronos)
[![Latest Release](https://img.shields.io/github/v/release/taciomcosta/kronos?include_prereleases)](https://github.com/taciomcosta/kronos/releases/latest)

**This project is being developed using [RDD](https://tom.preston-werner.com/2010/08/23/readme-driven-development.html) and is not production-ready yet!**
<br>Check the [architecture](docs/ARCHITECTURE.md) of the project and feel free to contribute. :)

## Table of Contents
- [Installation](#Installation)
  - [GitHub](#GitHub)
  - [Linux](#Linux)
  - [macOS](#macOS)
- [Commands](#Commands)
- [Examples](#Examples)
  - [Jobs](#Jobs)
  - [Notifier](#Notifiers)
    - [Slack](#Slack)


## Installation
### GitHub
You can find available binaries [here](https://github.com/taciomcosta/kronos/releases).

### Linux

**Arch**
```
git clone https://github.com/taciomcosta/kronos
makepkg -Si
```

### macOs
```
brew tap taciomcosta/kronos
brew install kronos
brew services start kronos

// Removing: 
brew services stop kronos
brew uninstall kronos
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
  create         Creates a new job/notifier
  delete         Deletes a job/notifier
  list           Lists all jobs/notifiers
  describe       Shows detailed information about a job/notifier
  history        Shows execution history of a job
  enable         Enables a job execution
  disable        Disables a job execution
  assign         Assigns a notifier to a job
  unassign       Unassign a notifier from a job

Flags:
  -h, --help   help for kronos

Use "kronos <command> --help" to learn more about a specific command.
```

## Examples

### Jobs

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
 
Assigned Notifiers:
  - my-slack
```

Showing execution history of all jobs:
```
> kronos history
NAME            EXECUTION                  STATUS           CPU TIME (ns)        MEM USAGE (MB)
myjob           2021-01-01 00:00:00        Succeeded        12320                364
myfiles         2021-01-02 00:00:00        Failed           7513                 60
backup-db       2021-01-03 00:00:00        Succeeded        6470                 211
```

Paginating executions of a specific job:
```
> kronos history myjob --page 3
NAME            EXECUTION                  STATUS           CPU TIME (ns)        MEM USAGE (MB)
myjob           2021-01-01 00:00:00        Succeeded        12320                364
myjob           2021-01-02 00:00:00        Succeeded        12320                364
myjob           2021-01-03 00:00:00        Failed           0                    0
```


Deleting a job:
```
> kronos delete job myjob
```

### Notifiers

#### Slack

Add kronos's slack bot to your workspace <br>
<a href="https://slack.com/oauth/v2/authorize?client_id=1880148272661.1937554624689&scope=chat:write&user_scope="><img alt="Add to Slack" height="40" width="139" src="https://platform.slack-edge.com/img/add_to_slack.png" srcSet="https://platform.slack-edge.com/img/add_to_slack.png 1x, https://platform.slack-edge.com/img/add_to_slack@2x.png 2x" /></a>

Then you can create a slack notifier with the auth token generated and a comma separated list of slack channel ids:
``` 
> kronos create notifier slack --name myslack --channel-ids 123456,123456 --auth-token #123
myslack created successfully
```

Lastly, you can assign your new slack notifier to any job you want
```
> kronos assign myslack myjob
myslack assigned to myjob

> kronos assign myslack myjob2 --errors-only
myslack assigned to myjob2
```

Or unassign it:
```
> kronos unassign myslack myjob
myslack unassigned from myjob
```
