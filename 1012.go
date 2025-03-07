package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, atr, ac, at, aq, ar float64
	fmt.Scan(&a, &b, &c)
	atr = (a * c) / 2
	ac = math.Pi * math.Pow(c, 2)
	at = ((a + b) * c) / 2
	aq = math.Pow(b, 2)
	ar = a * b
	fmt.Printf("TRIANGULO: %.3f\n", atr)
	fmt.Printf("CIRCULO: %.3f\n", ac)
	fmt.Printf("TRAPEZIO: %.3f\n", at)
	fmt.Printf("QUADRADO: %.3f\n", aq)
	fmt.Printf("RETANGULO: %.3f\n", ar)
}
