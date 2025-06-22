/*
Copyright © 2025 github.com/sush-il
*/

package cmd

import (
	"github.com/spf13/cobra"
	jsonUtils "github.com/sush-il/devNeeds/plugins/json"
)

var( 
	formatFileFlag bool
	checkValidFlag bool
)

var jsonUtilsCmd = &cobra.Command{
	Use:   "json",
	Short: "Various utilities to work with json files/strings",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if formatFileFlag {
			jsonUtils.FormatJSON(args)
		} else if checkValidFlag {
			jsonUtils.CheckValidJSON(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(jsonUtilsCmd)

	// flags and configuration settings.
	jsonUtilsCmd.Flags().BoolVarP(&formatFileFlag, "format", "f", false, "Pretty format a JSON file")
	jsonUtilsCmd.Flags().BoolVarP(&checkValidFlag, "checkValid", "c", false, "Check for errors in JSON syntax")

}
