/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/seungjinyu/gitcode/phase"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show commands shows the github repo on your browser immediately",
	Long:  `Show uses the .git folder and finds out the url and opens it on the browser`,
	Run: func(cmd *cobra.Command, args []string) {
		phase.Show(args)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
