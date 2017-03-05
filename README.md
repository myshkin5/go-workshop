# go-workshop
Material for a Go workshop

## Prerequisites
Installing the following items prior to the workshop will help get things moving more quickly.

1. Go tools <https://golang.org/doc/install>: Includes the Go compiler and other required tools.
    ```
    brew install go
    ```

1. Go Plugin for IntelliJ <https://github.com/go-lang-plugin-org> (Optional): Syntax highlighting and language insights
    1. Preferences -> Plugins -> Go
    1. Restart IntelliJ

1. `direnv` <https://direnv.net/> (Optional): Handy tool for setting environment variable per directory.
    ```
    brew install direnv
    ```
    Note `direnv` simply sets environment variables defined in `./.envrc` when entering a directory (or its children directories) and unsets the environment variables when `cd`'ing out. Look for `./.envrc` files and set the variables manually as an alternative.

## Introduction

1. [Hello World!](./hello-world/README.md)
2. [What is Go](./what-is-go/README.md)
