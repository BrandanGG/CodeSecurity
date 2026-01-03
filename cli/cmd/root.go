package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tyr",
	Short: "Tyr is a CLI tool for securing your code",
	Long:  `Tyr is a CLI tool for securiing your code through Static Analysis and Composition Analysis`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if version flag is set
		if showVersion, _ := cmd.Flags().GetBool("version"); showVersion {
			printVersion()
			return
		} else if scanPath, _ := cmd.Flags().GetString("path"); scanPath != "" {
			scanCmd.Run(cmd, args)
			return
		} else {
			// If no args and no flags, show help
			cmd.Help()
		}
	},
}

// execute the root command
func Execute() {
	rootCmd.Execute()
}
