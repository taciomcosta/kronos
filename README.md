<p align="center">
  <img src="./docs/kronos.png">
  <br><br>
  Kronos is a cross-platform job scheduler that helps you manage, monitor and inspect jobs.
</p>


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

Use "kronos <command> --help" to learn more about a specific command.
```

## Examples

Creating a new job with sugar syntax for interval: 
```
> kronos create job --name myjob --command ./my-job.sh --every-day
```
Alternatively, we can do:
```
> kronos create job --name myjob --command ./my-job.sh --interval "0 0 */1 * *"
```

Listing jobs:
```
> kronos list
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

