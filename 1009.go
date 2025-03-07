package main

import (
	"fmt"
)

func main() {
	var salary, sales, total float32
	var name string

	fmt.Scan(&name, &salary, &sales)
	total = salary + (sales * 0.15)
	fmt.Printf("TOTAL = R$ %.2f\n", total)
}
