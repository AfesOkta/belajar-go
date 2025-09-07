package app

import "fmt"

func HandlerMap() {
	harga := map[string]int{
		"Apel":   5000,
		"Mangga": 7000,
		"Jeruk":  6000,
	}

	fmt.Println("Harga Apel:", harga["Apel"])
	fmt.Println("Harga Mangga:", harga["Mangga"])
}
