package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var count int

	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		var n1, d1, n2, d2, num, den int
		var op rune
		fmt.Scanf("%d / %d %c %d / %d", &n1, &d1, &op, &n2, &d2)
		switch op {
		case '+':
			num = n1*d2 + n2*d1
			den = d1 * d2
		case '-':
			num = n1*d2 - n2*d1
			den = d1 * d2
		case '*':
			num = n1 * n2
			den = d1 * d2
		case '/':
			num = n1 * d2
			den = n2 * d1
		}
		divisor := mdc(abs(num), abs(den))
		simplifiedNum := num / divisor
		simplifiedDen := den / divisor

		if simplifiedDen < 0 {
			simplifiedNum = -simplifiedNum
			simplifiedDen = -simplifiedDen
		}
		fmt.Printf("%d/%d = %d/%d\n", num, den, simplifiedNum, simplifiedDen)
	}
}
