# Docker Dynamic Composer (WIP)

[![Build Status](https://travis-ci.org/MarioCarrion/dynamic-docker-composer.svg?branch=master)](https://travis-ci.org/MarioCarrion/dynamic-docker-composer) [![Code Climate](https://codeclimate.com/github/MarioCarrion/dynamic-docker-composer/badges/gpa.svg)](https://codeclimate.com/github/MarioCarrion/dynamic-docker-composer)

Project for creating dynamic composer files after a sequence of commands.

## Workflow

```
ddc-mc init --name example                                # Initiates the "example" workspace
ddc-mc template --add ./template1 --name template1        # Adds a template called "template1" from file "./template1"
ddc-mc template --add ./templateA --name templateA        # Adds a template called "templateA" from file "./templateA"
ddc-mc add --template template1 --container container-123 # Adds "container-123" using the template "template1"
ddc-mc add --template templateA --container container-abc # Adds "container-abc" using the template "templateA"
ddc-mc write                                              # Writes a docker compose file called "example.xml"
```

With that `docker-compose up --file example.yml` can be used

## Build

```go build -o ddc-mc```
