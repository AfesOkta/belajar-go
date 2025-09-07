package app

import (
	"fmt"
)

func HandlerSlice() {
	var hobis = []string{"Membaca", "Menulis", "Belanja"}
	hobis = append(hobis, "Berenang", "Ngegame")

	for i, hobi := range hobis {
		fmt.Println(i, hobi)
	}
}
