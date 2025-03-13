package main

import (
	"discoverypolling/server"
	"discoverypolling/utils/logger"
	"log"
)

func main() {
	log.Printf("Starting discovery-polling...")

	if err := server.RunServer("tcp://127.0.0.1:8080"); err != nil {
		logger.Error("Error starting server:", err)
	}
}
