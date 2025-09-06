package app

import "fmt"

type HelloWorldService struct{}

func NewHelloWorldService() *HelloWorldService {
	return &HelloWorldService{}
}

// Method untuk generate pesan
func (s *HelloWorldService) GetMessage(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to Golang 🚀", name)
}
