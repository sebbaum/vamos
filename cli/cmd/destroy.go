
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy the local development environment and all corresponding files.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("destroy called. However it is not yet implemented.")
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}
