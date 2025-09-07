package app

import (
	"fmt"
)

func HandlerArray() {
	var angka [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println(angka[0]) // 10
	fmt.Println(angka[1]) // 20
	fmt.Println(angka[2]) // 30
	fmt.Println(angka[3]) // 40
	fmt.Println(angka[4]) // 50
}
