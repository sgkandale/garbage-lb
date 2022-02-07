package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "garbagelb",
	Short: "Garbage Load Balancer",
	Long:  `A simple load balancer with a web interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		operator()
	},
}

func Execute() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
