package codec

import (
	"fmt"
	"log"
	"net/url"
)

func EncodeURL(urls []string) {
	var encodedURLs []string

	log.Println("Starting to strings into URL format.")

	for _, rawURL := range urls {
		encodedURL := url.PathEscape(rawURL)
		encodedURLs = append(encodedURLs, encodedURL)
		fmt.Println(encodedURL)
	}
}
