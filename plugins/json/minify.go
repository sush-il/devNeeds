package jsonUtils

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/sush-il/devNeeds/utils/pluginsCommon"
)

func Minify(jsonFilePaths []string) {
	log.Println("Starting to minify the JSON files.")
	for _, filePath := range jsonFilePaths {
		fileContent, err := pluginsCommon.ReadFile(filePath)

		if err != nil {
			log.Println(err)
			continue
		}

		var compactContents bytes.Buffer
		if err := json.Compact(&compactContents, fileContent); err != nil {
			log.Printf("Error occurred while compacting JSON: %v", err)
			continue
		}

		writeFileErr := pluginsCommon.WriteFile(filePath, compactContents.Bytes())
		if writeFileErr != nil {
			log.Printf("Error occurred while writing compact version to %s. %v", filePath, err)
			continue
		}

		log.Println("Completed creating a compact version of JSON files.")

	}
}
