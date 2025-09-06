package main

import (
	"fmt"
	"hello-world/internal/app"
)

func main() {
	service := app.NewHelloWorldService()
	message := service.GetMessage("Afes")
	fmt.Println(message)
}
