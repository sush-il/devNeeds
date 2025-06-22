package jsonUtils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func FormatJSON(jsonFilePaths []string) string{
	fmt.Println("Formatting JSON")
	jsonStyleString := `{"hello":{"someKey": "someValue"}}`
	var buf bytes.Buffer
	if err := json.Indent(&buf, []byte(jsonStyleString), "", "  "); err != nil {
		fmt.Println(buf.String());
	}
	return "hello"
}