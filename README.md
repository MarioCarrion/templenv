# templenv

## Install

`go get github.com/MarioCarrion/templenv`

## What is it?

**templenv** is a command line program for generating output from environment variables by using Go templates, a key piece in the **Docker Full Development Workflow**. `os.StdOut` is used for the writer when executing the template.

## Templates

**templenv** uses [text/template](https://golang.org/pkg/text/template/) for generating the final file. Please review the official documentation for learning how to write those templates.

## Environment Variables

**templenv** uses environment variables for dynamically writing the output file. Any environment variable can be accessed from within **templenv** though the internal template function **getEnv**.

## Docker

You can use the included `Dockerfile` for building your own image, or if you prefer a smaller one you can use the `Dockerfile.static`, see [DockerHub](https://hub.docker.com/r/mariocarrion/templenv/) as well.

# Example

Assuming you have a template `snowman.tmpl` with the following content:

```
My name is {{ getEnv "USER" }} and I like warm hugs!
```

and your current username is `mario`, by using `templenv snowman.tmpl` the following output will be generated:

```
My name is mario and I like warm hugs!
```
