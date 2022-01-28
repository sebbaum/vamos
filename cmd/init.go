/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new vamos.yaml file for your environment",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called. However it is not yet implemented")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
