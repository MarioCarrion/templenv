# tempenv (WIP)

## What is it?

**tempenv** is a command line program used for generating files using a combination of Go templates and environment variables, a key piece in the **Docker Quality Assurance Workflow**.

# Templates

**tempenv** uses [text/template](https://golang.org/pkg/text/template/) for generating the final file. Please review the official documentation for learning how to write those templates.

# Environment Variables

**tempenv** uses environment variables for dynamically writing the output file. Those variables are supposed to be prefixed by a specific string to avoid collisions with other environment variables, the default prefix is **TENV**. There are two ways to define the environment variables: either by using the _live environment_ or by loading an _environment file_.

## Live environment

Just export the variables when calling the **tempenv** program to allow it to have access to those variables. For example:

```
DCC_HELLO="hello world" mc-tempenv --template ./docker-compose.tmpl --output ./docker-compose.yml
```

## Environment file



loads the environment variables from either the live environment or from a file, it allows you to specify a _prefix_; this prefix is meant to be used for avoiding collisions with other environment variables for example you could define a variable .

```
dcc-mc --template docker-compose.tmpl --output docker-compose-something.yml --env_file env.dev --env_prefix xyz
```
With that `docker-compose up --file example.yml` can be used

## Build

```go build -o ddc-mc```
