package main

import "fmt"

func main() {
	var id, quant int
	var price, total float32

	total = 0
	for i := 0; i < 2; i++ {
		fmt.Scan(&id, &quant, &price)
		total = total + (float32(quant) * price)
		price = 0
		quant = 0
	}
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
