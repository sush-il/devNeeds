package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sush-il/devNeeds/plugins/codec"
)

var (
	encodeURL      bool
	converToBase64 bool
)

var codecCmd = &cobra.Command{
	Use:   "codec",
	Short: "Encoding & decoding of strings and files.",
	Run: func(cmd *cobra.Command, args []string) {
		if encodeURL {
			codec.EncodeURL(args)
		} else if converToBase64 {
			codec.ConvertToBase64(args)
		}
	},
}

func init() {

	rootCmd.AddCommand(codecCmd)

	codecCmd.Flags().BoolVar(&encodeURL, "encodeURL", false, "Encode URL")
	codecCmd.Flags().BoolVarP(&converToBase64, "converToBase64", "b", false, "Encode text file to base64")
}
