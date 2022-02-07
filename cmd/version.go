package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of GarbageLB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GarbageLB v0.1.0")
	},
}
