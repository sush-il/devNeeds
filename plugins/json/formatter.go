package jsonUtils

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/sush-il/devNeeds/utils/pluginsCommon"
)

func FormatJSON(jsonFilePaths []string, indentType string, spaceIndentLevel int) {
	log.Println("Formatting of JSON file/s started.")
	
	var indent string
	if (indentType == "space") {
		indent = strings.Repeat(" ", spaceIndentLevel) 
	} else if (indentType == "tab"){
		indent = "\t"
	}
	
	for _, filePath := range jsonFilePaths {
		fileContent, err := pluginsCommon.ReadFile(filePath)

		if err != nil {
			log.Println(err)
			continue	
		}
		var formattedFileContent bytes.Buffer
		if err := json.Indent(&formattedFileContent, []byte(fileContent), "", indent); err != nil {
			log.Printf("Failed when formatting %s -> %v\n",filePath, err)
			continue	
		}

		pluginsCommon.WriteFile(filePath, formattedFileContent.Bytes())
	}
	log.Println("Formatting of JSON file/s is complete.")
}