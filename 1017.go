package main

import "fmt"

func main() {
	var time, speed int
	fmt.Scan(&time, &speed)
	space := time * speed
	gas := float64(space) / 12.0
	fmt.Printf("%.3f\n", gas)
}
