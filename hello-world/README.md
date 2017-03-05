# Hello World!

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

- The `main` package
    - Special package for creating an executable
    - More on packages later
- Imports
    - Pull in Go code from other packages
    - Alias mechanism
- Functions
    - `main()` function - like a static method in Java
    - Calls `fmt.Println()` function
- Execution
    Requires a `main` package with a `main()` function that takes no arguments and returns no value
    - `go run hello-world.go`
    - `go build hello-world.go`
        ```bash
        $ ls -lh hello-world
        -rwxr-xr-x  1 dwayne  staff   1.6M Mar  4 22:18 hello-world
        ```

Next [What is Go](../what-is-go/README.md)
