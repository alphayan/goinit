package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of goinit",
	Long:  `All software has versions. This is goinit's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("goinit version v0.1.0")
	},
}
