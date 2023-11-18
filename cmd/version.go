package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints out the application version.",
	Run: func(cmd *cobra.Command, args []string) {
		version := os.Getenv("APPLICATION_VERSION")
		if version == "" {
			version = "Unkown"
		}

		fmt.Printf("Application version: %s\n", version)
		fmt.Printf("Go version: %s\n", runtime.Version())
		fmt.Printf("Architecture: %s\n", runtime.GOARCH)
		fmt.Printf("Operating system: %s\n", runtime.GOOS)
	},
}
