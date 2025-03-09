package main

import "fmt"

func main() {
	var secs, hour, min int
	fmt.Scan(&secs)

	for secs >= 3600 {
		secs -= 3600
		hour++
	}
	for secs >= 60 {
		secs -= 60
		min++
	}

	fmt.Printf("%d:%d:%d\n", hour, min, secs)
}
