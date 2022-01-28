package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start up the local development environment",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("up called. However it in not yet implemented.")
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
