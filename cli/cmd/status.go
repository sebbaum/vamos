
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of your local development stack.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status called. Hever it is not yet implemented.")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
