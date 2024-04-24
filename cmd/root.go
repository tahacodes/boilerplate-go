package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "The root command of the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Visit: t.ly/xzLR")
	},
}

func Execute() {
	// Disable automatic generation of the help subcommand
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
