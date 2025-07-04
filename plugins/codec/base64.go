package codec

import (
	"encoding/base64"
	"log"

	"github.com/sush-il/devNeeds/utils/pluginsCommon"
)

func ConvertToBase64(paths []string) {

	for _, path := range paths {
		fileContent, err := pluginsCommon.ReadFile(path)

		if err != nil {
			log.Print(err)
		}

		encodedText := base64.StdEncoding.EncodeToString(fileContent)
		pluginsCommon.WriteFile(path, []byte(encodedText))
	}

}
