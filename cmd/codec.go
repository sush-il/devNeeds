package cmd

import (
	"github.com/spf13/cobra"
	encode "github.com/sush-il/devNeeds/plugins/codec"
)

var(
	encodeURL bool
)

var codecCmd = &cobra.Command{
	Use: "codec",
	Short: "Encoding & decoding of strings and files.",
	Run: func(cmd *cobra.Command, args []string) {
		if encodeURL {
			encode.EncodeURL(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(codecCmd)

	codecCmd.Flags().BoolVar(&encodeURL, "encodeURL", false, "Encode URL")
}
