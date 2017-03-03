# go-workshop
Material for a Go workshop

## Prerequisites
Installing the following items prior to the workshop will help get things moving more quickly.

1. Go tools <https://golang.org/doc/install>: Includes the Go compiler and other required tools.
    ```
    brew install go
    ```

1. `direnv` <https://direnv.net/> (Optional): Handy tool for setting environment variable per directory.
    ```
    brew instal direnv
    ```
    Note `direnv` simply sets environment variables defined in `./.envrc` when entering a directory (or its children directories) and unsets the environment variables when `cd`'ing out. Look for `./.envrc` files and set the variables manually as an alternative.
