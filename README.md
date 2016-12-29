# templenv

## Install

`go get github.com/MarioCarrion/templenv`

## What is it?

**templenv** is a command line program for generating output from environment variables by using Go templates, a key piece in the **Docker Full Development Workflow**. `os.StdOut` is used for the writer when executing the template.

# Templates

**templenv** uses [text/template](https://golang.org/pkg/text/template/) for generating the final file. Please review the official documentation for learning how to write those templates.

# Environment Variables

**templenv** uses environment variables for dynamically writing the output file. Any environment variable can be accessed from within **templenv** though the internal template function **getEnv**.
