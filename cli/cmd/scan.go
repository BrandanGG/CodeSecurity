package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	scanPath  string
	scanTypes []string
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the code for vulnerabilities",
	Long:  `Scan the code for vulnerabilities using Static Analysis and Composition Analysis`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Scanning code at %s\n", scanPath)
		fmt.Printf("Executing scan type(s): %s\n", strings.Join(scanTypes, ", "))
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	rootCmd.PersistentFlags().StringVarP(&scanPath, "path", "p", "", "Path to the code to scan")
	rootCmd.PersistentFlags().StringSliceVarP(&scanTypes, "types", "t", []string{"all"}, "Types of scans to run")
}
