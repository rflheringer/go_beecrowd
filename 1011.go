package main

import (
	"fmt"
	"math"
)

func main() {
	var r, vol float64
	fmt.Scan(&r)
	vol = (4.0 / 3.0) * math.Pi * math.Pow(r, 3)
	fmt.Printf("VOLUME = %.3f\n", vol)
}
