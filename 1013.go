package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)
	maiorAB := (a + b + int(math.Abs(float64(a-b)))) / 2
	maior := (maiorAB + c + int(math.Abs(float64(maiorAB-c)))) / 2
	fmt.Printf("%d eh o maior\n", maior)
}
