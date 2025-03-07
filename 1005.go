package main

import "fmt"

func main() {
	var x, y, m float32

	fmt.Scan(&x, &y)
	m = ((x * 3.5) + (y * 7.5)) / 11
	fmt.Printf("MEDIA = %.5f\n", m)
}
