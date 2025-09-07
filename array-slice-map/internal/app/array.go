package app

import "fmt"

func HandlerArray() {
	var angka [3]int = [3]int{10, 20, 30}
	fmt.Println(angka[0]) // 10
	fmt.Println(angka[1]) // 20
	fmt.Println(angka[2]) // 30
}
