package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build all Dockerfiles, yaml and configs.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called. However it is not yet implemented.")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
