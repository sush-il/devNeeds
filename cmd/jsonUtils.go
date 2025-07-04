/*
Copyright © 2025 github.com/sush-il
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	jsonUtils "github.com/sush-il/devNeeds/plugins/json"
)

var (
	formatFileFlag bool
	checkValidFlag bool
	minifyFlag     bool
	indentType     string
	indentLevel    int
)

var jsonUtilsCmd = &cobra.Command{
	Use:   "json",
	Short: "Various utilities to work with json files/strings",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
	
		if err:= validateFlags(cmd); err != nil {
			return err
		}	
		
		if formatFileFlag {
			jsonUtils.FormatJSON(args, indentType, indentLevel)
		} else if checkValidFlag {
			jsonUtils.CheckValidJSON(args)
		} else if minifyFlag {
			jsonUtils.Minify(args)
		}

		return nil
	},
}

func validateFlags(cmd *cobra.Command) error {
	if cmd.Flags().Changed("indentType") && !formatFileFlag {
		return fmt.Errorf("indentType flag can only be used with format flag")	
	}

	if cmd.Flags().Changed("indentLevel") && indentType != "space" {
		return fmt.Errorf("indentLevel can only be set when indentType is 'space'")	
	}

	actions := 0
	if formatFileFlag {
		actions++
	}
	if checkValidFlag {
		actions++
	}
	if minifyFlag {
		actions++
	}

	if actions == 0 {
		return fmt.Errorf("Please provide one of --format, --checkValid, or --minify")
	}
	if actions > 1 {
		return fmt.Errorf("Please use only one of --format, --checkValid, or --minify at a time")
	}

	return nil
}

func init() {
	rootCmd.AddCommand(jsonUtilsCmd)

	// flags and configuration settings.
	jsonUtilsCmd.Flags().BoolVarP(&formatFileFlag, "format", "f", false, "Pretty format a JSON file")
	jsonUtilsCmd.Flags().BoolVarP(&checkValidFlag, "checkValid", "v", false, "Check for errors in JSON syntax")
	jsonUtilsCmd.Flags().StringVar(&indentType, "indentType", "space", "Type of indent: space or tab")
	jsonUtilsCmd.Flags().IntVar(&indentLevel, "indentLevel", 2, "Indent Level when indent type of space")
	jsonUtilsCmd.Flags().BoolVarP(&minifyFlag, "minify", "m", false, "Remove whitespace from JSON files.")

}
