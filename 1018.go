package main

import "fmt"

func main() {
	var money, count int
	fmt.Scan(&money)
	fmt.Printf("%d\n", money)
	for money >= 100 {
		money = money - 100
		count++
	}
	fmt.Printf("%d nota(s) de R$ 100,00\n", count)
	count = 0
	for money >= 50 {
		money = money - 50
		count++
	}
	fmt.Printf("%d nota(s) de R$ 50,00\n", count)
	count = 0
	for money >= 20 {
		money = money - 20
		count++
	}
	fmt.Printf("%d nota(s) de R$ 20,00\n", count)
	count = 0
	for money >= 10 {
		money = money - 10
		count++
	}
	fmt.Printf("%d nota(s) de R$ 10,00\n", count)
	count = 0
	for money >= 5 {
		money = money - 5
		count++
	}
	fmt.Printf("%d nota(s) de R$ 5,00\n", count)
	count = 0
	for money >= 2 {
		money = money - 2
		count++
	}
	fmt.Printf("%d nota(s) de R$ 2,00\n", count)
	count = 0
	for money >= 1 {
		money = money - 1
		count++
	}
	fmt.Printf("%d nota(s) de R$ 1,00\n", count)
	count = 0
}
