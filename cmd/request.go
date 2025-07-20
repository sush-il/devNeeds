package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	request "github.com/sush-il/devNeeds/plugins/request"
)

var (
	url        string
	headers    string
	body       string
	getMethod  bool
	postMethod bool
	method     string
)

var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "Make requests to APIs",
	Args:  cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		validateRequestFlags()
		request.MakeRequest(url, method, headers, body)

		if err := validateRequestFlags(); err != nil {
			return err
		}

		return nil
	},
}

func validateRequestFlags() error {
	if url == "" {
		log.Println("URL is required to make an API call")
		return fmt.Errorf("URL not provided")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(requestCmd)

	requestCmd.Flags().StringVarP(&url, "url", "u", "", "API URL")
	requestCmd.Flags().StringVarP(&headers, "headers", "h", "", "Request headers")
	requestCmd.Flags().StringVarP(&body, "body", "b", "", "Request body")
}
