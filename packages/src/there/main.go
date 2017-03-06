package main

import (
	"fmt"
	abc "a/b/c/stuff"
	xyz "x/y/z/stuff"
)

func main() {
	fmt.Printf("Here's the int: %d and here's the string: %s\n", abc.PublicInt(), xyz.PublicString())
}
