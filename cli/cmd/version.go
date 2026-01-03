package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	showVersion bool
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Tyr",
	Long:  `Print the version number, commit hash, and build time of Tyr.`,
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	fmt.Printf("Tyr version 0.0.1\n")
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Print the version number of Tyr")
}
