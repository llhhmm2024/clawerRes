package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "v0.0.1"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of crawler",
	Long:  "All software has versions. This is crawler's",
	Run: func(cmd *cobra.Command, args []string) {
		GetVersion()
	},
}

func GetVersion() {
	fmt.Println("crawler version: ", version)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
