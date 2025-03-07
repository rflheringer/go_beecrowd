package main

import "fmt"

func main() {
	var km int
	var gas float64
	fmt.Scan(&km, &gas)
	total := float64(km) / gas
	fmt.Printf("%.3f km/l\n", total)
}
