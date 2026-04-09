package main

import (
	"fmt"
	"log"
	"os"
)

// middleware-chain — Composable middleware pipeline with context propagation and error recovery
func main() {
	logger := log.New(os.Stdout, "[middleware-chain] ", log.LstdFlags)
	logger.Println("Starting application...")

	if err := run(); err != nil {
		logger.Fatalf("Fatal error: %v", err)
	}
}

func run() error {
	fmt.Println("Application initialized successfully")
	return nil
}
