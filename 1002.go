package main

import (
	"fmt"
	"math"
)

func main() {
	var x float32
	fmt.Scan(&x)
	fmt.Printf("A=%.4f\n", math.Pow(x, 2)*math.Pi)
}
