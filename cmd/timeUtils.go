/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	timeUtils "github.com/sush-il/devNeeds/plugins/time"
)

var (
	timestamp string
)

var timeUtilsCmd = &cobra.Command{
	Use:   "timeUtils",
	Short: "Various options to format / generate  time functions",
	Run: func(cmd *cobra.Command, args []string) {
		timeUtils.ConvertToUnix(args)	
	},
}

func init() {
	rootCmd.AddCommand(timeUtilsCmd)

	timeUtilsCmd.Flags().StringVarP(&timestamp, "unix", "u", "", "Convert an ISO timestamp string to Unix format")
}
