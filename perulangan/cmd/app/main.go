package main

import (
	"fmt"
)

func main() {
	// for biasa
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	// for sebagai while
	x := 1
	for x <= 3 {
		fmt.Println("x =", x)
		x++
	}
}
