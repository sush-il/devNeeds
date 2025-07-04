package jsonUtils

import (
	"encoding/json"
	"log"

	"github.com/sush-il/devNeeds/utils/pluginsCommon"
)

func CheckValidJSON(jsonFilePaths []string) {
	for _, filePath := range jsonFilePaths {
		fileContent, err := pluginsCommon.ReadFile(filePath)
		if err != nil {
			log.Println(err)
			continue
		}
		var unmarshalValue any
		unmarshalError := json.Unmarshal(fileContent, &unmarshalValue)

		if unmarshalError != nil {
			log.Printf("The file %s is invalid; %v", filePath, unmarshalError)
			continue
		}

		log.Printf("The file %s is valid.", filePath)
	}
	log.Println("Validation of all JSON files is complete.")
}
