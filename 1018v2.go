package main

import "fmt"

func main() {
	var money int
	fmt.Scan(&money)

	notes := []int{100, 50, 20, 10, 5, 2, 1}
	for _, note := range notes {
		count := money / note
		money %= note
		fmt.Printf("%d nota(s) de R$ %d,00\n", count, note)
	}
}
