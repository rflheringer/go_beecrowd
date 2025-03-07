package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)
	x := math.Pow((x2 - x1), 2)
	y := math.Pow((y2 - y1), 2)
	result := math.Sqrt((x + y))
	fmt.Printf("%.4f\n", result)
}
