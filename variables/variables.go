package main

import (
	"fmt"
	"math"
)

func main() {

	// Simple variable declaration -- has zero value
	var myInt int
	var myString string
	var myBool bool
	var myFloat float64

	fmt.Printf(`These are the zero values for an int %d, a string "%s", a bool %v and a float %v`,
		myInt, myString, myBool, myFloat)
	fmt.Println()

	// There is type inference for "short assignments. This will have a type of `int`
	anotherInt := 3
	fmt.Printf("Initial value %d\n", anotherInt)

	// Now you can't assign something of another type
	//anotherInt = "hi"

	// But you can assign another int
	anotherInt = 6
	fmt.Printf("Reassigned value %d\n", anotherInt)

	// Constants can have a crazy amount of precision until they get assigned and have a type
	pi := math.Pi
	fmt.Printf("pi = %v and is is of type %T\n", pi, pi)

	// All conversions are explicit (more on that later) via type conversions (like a static cast)
	three := int(pi)
	fmt.Printf("pi with a type conversion to int is %v\n", three)

	// Complete list of basic types
	// bool

	// string

	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr

	// byte // alias for uint8

	// rune // alias for int32
	// represents a Unicode code point

	// float32 float64

	// complex64 complex128
}
