package main

import "fmt"

func main() {
	var id, hour int
	var salary, final float32

	fmt.Scan(&id, &hour, &salary)
	final = salary * float32(hour)
	fmt.Printf("NUMBER = %d\n", id)
	fmt.Printf("SALARY = U$ %.2f\n", final)
}
