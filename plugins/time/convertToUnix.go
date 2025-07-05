package timeUtils

import (
	"fmt"
	"log"
	"time"
)

func ConvertToUnix(timestamps []string) ([]int64, error) {
	
	var unixTimestamps []int64

	for _,timestamp := range timestamps {
		t, err := time.Parse(time.RFC3339, timestamp)
		
		if err != nil {
			return nil,err
		}

		fmt.Printf("Time stamp %s converted to %d\n", timestamp, t.Unix())
		unixTimestamps = append(unixTimestamps, t.Unix())
	}

	log.Println("Completed converting timestamps to unix format")
	
	return unixTimestamps, nil
}
