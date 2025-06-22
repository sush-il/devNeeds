package jsonUtils

import (
	"encoding/json"
	"fmt"
)

func CheckValidJSON(jsonFilePaths []string) string{

	jsonString := `{"hello":{"someKey": "someValue"}}`
    var unmarshalledValue any

    err := json.Unmarshal([]byte(jsonString), &unmarshalledValue)
    if err != nil {
        fmt.Println("Invalid json.", err) 
		return fmt.Sprintf("Invalid JSON: %v", err)
    }

    fmt.Println("Parsed JSON:", unmarshalledValue)
    return "JSON is valid"
}