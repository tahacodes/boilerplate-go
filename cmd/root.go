package cmd

import (
	"log"

	"github.com/tahacodes/go-boilerplate/configs"
)

func Execute() {
	// Shutdown gracefully.
	defer close()

	// Starting point of the application
	log.Println("Starting the application...")

	// Initiate application configs
	if err := configs.InitConfigs(); err != nil {
		log.Fatalf("failed to initiate application configs, %v", err)
	}
}

func close() {
	// Terminating the application. Here's some recommendations
	// Deregister your app from service discovery if needed
	// Stop accepting or executing new jobs or requests
	// Close any stateful databases and external connections
	// Allow currently running requests, jobs or goroutines to complete their process.
	log.Printf("Shutting down...")
}
