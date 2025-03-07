package main

import "fmt"

func main() {
	var x, y, z, m float32

	fmt.Scan(&x, &y, &z)
	m = (x*2 + y*3 + z*5) / 10
	fmt.Printf("MEDIA = %.1f\n", m)
}
