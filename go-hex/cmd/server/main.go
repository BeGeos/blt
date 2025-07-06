package main

import (
	"go-hex/internal/adapter/http/echo"
	"log"
)

func main() {
	if err := echo.NewServer().Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
