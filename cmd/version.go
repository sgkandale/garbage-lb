package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "v0.1.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of GarbageLB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GarbageLB version:", Version)
		os.Exit(0)
	},
}
