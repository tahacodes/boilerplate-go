package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts the API server.",
	Run: func(cmd *cobra.Command, args []string) {
		// Shutdown gracefully.
		defer close()

		// Implement the logic here.
		log.Println("Starting the application...")
	},
}

func close() {
	// Terminating the application. Here's some recommended steps:
	// Deregister from service discovery if needed.
	// Stop accepting or executing new jobs or requests.
	// Close any stateful connections like databases and external resources.
	// Allow currently running requests, jobs or goroutines to complete their processing.
	log.Printf("Shutting down...")
}
