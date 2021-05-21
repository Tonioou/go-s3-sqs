package main

import (
	"cmd/main.go/internal/config"
	"fmt"
)

func main() {
	config := config.NewConfig()
	fmt.Println("Access Key:", config.AccessKey)
	fmt.Println("Secret Key:", config.SecretKey)
}
