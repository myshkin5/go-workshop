# Packages/Imports/Accessibility

- `GOPATH`
    - Typically just points to a Go workspace directory
    - Occasionally multiple workspaces (colon separated) for system/project workspaces
    - Defaults to the `$HOME/go` directory
    - For IntelliJ set via Preferences -> Languages & Frameworks -> Go -> Go Libraries -> Project Libraries -> +
- The Go workspace
    - `./src`: Holds project Go files and third-party source
    - `./pkg`: Generated binaries (not checked in)
    - `./bin`: Generated executables (not checked in)
        - Add to `PATH`: `export PATH=${GOPATH//://bin:}/bin:$PATH`
- Create a package with public and private functions
    - `./src/a/b/c/stuff/look_here.go`
    - `./src/x/y/z/stuff/and_here.go`
- Public names start with a capital letter
- Private names start with lowercase
- Create a `main()` function that prints the return values of the public functions
- To avoid import namespace collisions, add an alias right after `import`

Prev [What is Go](../what-is-go/README.md) | Next 
