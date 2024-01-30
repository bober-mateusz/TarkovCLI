package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "TarkovCLI",
	Short: "TarkovCLI is a small project focused around quick and easy CLI access to tarkov data on-the-go",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	// Any initialization logic for the root command goes here
}
