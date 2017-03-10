package main

import (
	"fmt"
	"math"
	"time"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 0; i < 5; i++ {
		z = z - (math.Pow(z, 2) - x) / 2 * z
	}
	return z
}

func main() {
	start := time.Now()
	fmt.Println(Sqrt(2))
	fmt.Println(time.Since(start).Nanoseconds())
	start = time.Now()
	fmt.Println(math.Sqrt(2))
	fmt.Println(time.Since(start).Nanoseconds())
}
