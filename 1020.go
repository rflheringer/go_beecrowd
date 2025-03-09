package main

import "fmt"

func main() {
	var age_in_days, months, year int
	fmt.Scan(&age_in_days)

	for age_in_days >= 365 {
		age_in_days -= 365
		year++
	}

	for age_in_days >= 30 {
		age_in_days -= 30
		months++
	}

	fmt.Printf("%d ano(s)\n", year)
	fmt.Printf("%d mes(es)\n", months)
	fmt.Printf("%d dia(s)\n", age_in_days)
}
