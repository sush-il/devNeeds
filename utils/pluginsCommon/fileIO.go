package pluginsCommon

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(filePath string)([]byte, error){
	file, err := os.ReadFile(filePath);

	if err != nil {
		return nil, fmt.Errorf("an error occurred while reading %s. %w",filePath, err)
	}
	
	return file, nil
}

func WriteFile(filePath string, data []byte) error{
	err := os.WriteFile(filePath, data, 0644)

	if err != nil {
		return fmt.Errorf("an error occurred when writing to %s. %w",filePath, err)
	}
	
	log.Printf("Write to %s successful.\n", filePath)
	return nil
}