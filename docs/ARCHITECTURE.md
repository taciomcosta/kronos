# Architecture

This document aims to outline the decisions behind the high level and low level
design of Kronos. Even though it's a simple project, I felt that it was needed
to document these details to keep the project healthy in the long term and
make people feel welcome to contribute.

## Requirements

The main quality attributes of Kronos are:
- portability: support many OS platforms;
- extensibility: support many channels, task automation, etc;
- usability: avoid that users have mental overhead when managing their jobs;

## High Level Design

Kronos is divided into two programs:

### kronosd

*kronosd* is the core of Kronos. It is the daemon responsible for
scheduling jobs, monitor their executions, and notify channels when an event occurs.
*kronosd* exposes a REST API that serves `kronoscli`, but that can also
be used for task automation by users.

### kronoscli
*kronoscli* is our frontend. It's the command line tool that allows users to 
interact with the daemon. It's supposed to be a thin layer with as few business
logic as possible. Its main subcomponents are a REST client and its commands that
were created using [Cobra](https://github.com/spf13/cobra).

The entrypoint of each program can be found at `cmd/` folder.

## Low Level Design

![clean architecture](docs/clean_architecture-kronos.png)

Kronos low level design was inspired by [Clean Architecture]().
In a few words, the main goal is to separate business rules (use cases and entities)
from details (rest, cli, sqlite, OS), as shown above.

We have a few rules about dependencies:
- use cases depend on entities; entities don't depend on use cases;
- details depend on business rules; business rules don't depend on details;
- details don't depend on each other;

The low level design structure can be found at `internal/` folder

## Testing

We have two categories of tests: unit tests and acceptance tests. 
But how does it fit into the design above?

- unit tests: unit tests should be placed at the edge of use cases. We want it
to test business rules despite if they're implemented in use cases or entities.
Why?
Because that makes our tests [more like functional tests than micro tests](https://blog.cleancoder.com/uncle-bob/2017/05/05/TestDefinitions.html)
and helps us to avoid a mockist style of testing things.

- acceptance tests: acceptance tests are placed at the edge of *kronosd*.
They are always written based on the users point of view. 
While unit tests target individual use cases and don't integrate with details (database, api, etc), 
acceptance tests generally test the combination of many usecases and can be thought as a special
case of integration test as well.

