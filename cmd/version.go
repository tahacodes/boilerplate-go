package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/tahacodes/go-boilerplate/configs"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints out the application version and runtime specifications.",
	Run: func(cmd *cobra.Command, args []string) {
		// Initiate application configs
		if err := configs.InitConfigs(); err != nil {
			log.Fatalf("failed to initiate application configs, %v", err)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
		fmt.Fprintf(w, "Version:\t%s\n", configs.Config.App.Version)
		fmt.Fprintf(w, "Go version:\t%s\n", runtime.Version())
		fmt.Fprintf(w, "Architecture:\t%s\n", runtime.GOARCH)
		fmt.Fprintf(w, "Operating system:\t%s\n", runtime.GOOS)
		w.Flush()
	},
}
