/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version",
	Long:  `Shows the gitcode version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitcode version :", os.Getenv("VERSION"))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
