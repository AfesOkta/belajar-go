package app

import "fmt"

func HandlerSlice() {
	buah := []string{"Apel", "Mangga", "Jeruk"}
	buah = append(buah, "Pisang") // nambah data

	for i, b := range buah {
		fmt.Println(i, b)
	}
}
