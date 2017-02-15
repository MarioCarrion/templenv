# templenv

[![Build Status](https://travis-ci.org/MarioCarrion/templenv.svg?branch=master)](https://travis-ci.org/MarioCarrion/templenv)

## Install

`go get github.com/MarioCarrion/templenv`

## What is it?

**templenv** is a command line program for generating output from environment variables by using Go templates, a key piece in the **Docker Full Development Workflow**. `os.StdOut` is used for the writer when executing the template.

## Templates

**templenv** uses [text/template](https://golang.org/pkg/text/template/) for generating the final file. Please review the official documentation for learning how to write those templates.

## Supported Functions

### getEnv

**getEnv** is a function for accessing environment variables from within **templenv**. For example, assuming you have a template `snowman.tmpl` with the following content:

```
My name is {{ getEnv "USER" }} and I like warm hugs!
```

and your current username is `mario`, by using `templenv snowman.tmpl` the following output will be generated:

```
My name is mario and I like warm hugs!
```

### loadEnvFilename

**loadEnvFilename** is a function for reading a file containing [environment variables](https://github.com/bkeepers/dotenv#usage) that then parses/converts them into what could look like Docker Compose [environment variables](https://docs.docker.com/compose/compose-file/#environment).

For example, assuming you have two files, the first one a template `docker-compose.tmpl` with the following content:

```
version: '2'
services:
  web:
    image: mariocarrion/templenv:latest
    command: /templenv
    environment:
{{ loadEnvFilename "      " "env.compose" }}
```

and the second one `env.compose` with the following content:

```
# Production env variables
HELLO=world!
```

The following output will be generated:

```
version: '2'
services:
  web:
    image: mariocarrion/templenv:latest
    command: /templenv
    environment:
      HELLO: world!
```

## Docker

You can use the included `Dockerfile` for building your own image, or if you prefer a smaller one you can use the `Dockerfile.static`, see [DockerHub](https://hub.docker.com/r/mariocarrion/templenv/) as well.
