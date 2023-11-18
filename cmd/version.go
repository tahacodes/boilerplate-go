package cmd

import (
	"fmt"
	"log"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/tahacodes/go-boilerplate/configs"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints out the application version.",
	Run: func(cmd *cobra.Command, args []string) {
		// Initiate application configs
		if err := configs.InitConfigs(); err != nil {
			log.Fatalf("failed to initiate application configs, %v", err)
		}

		fmt.Printf("Application version: %s\n", configs.Config.App.Version)
		fmt.Printf("Go version: %s\n", runtime.Version())
		fmt.Printf("Architecture: %s\n", runtime.GOARCH)
		fmt.Printf("Operating system: %s\n", runtime.GOOS)
	},
}
