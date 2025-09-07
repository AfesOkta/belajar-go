package main

import (
	"fmt"
)

func main() {
	score := 80

	if score >= 90 {
		fmt.Println("Nilai A")
	} else if score >= 75 {
		fmt.Println("Nilai B")
	} else {
		fmt.Println("Nilai C")
	}
}
