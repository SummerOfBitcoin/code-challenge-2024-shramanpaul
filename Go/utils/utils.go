package utils

import (
	"bufio"
	"log"
	"os"
)

func JsonData(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}


func writeToFile(TransactionIDs []string) {
	
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()

	
	bufferedWriter := bufio.NewWriter(file)

	// Writing the TransactionIDs to the file
	for _, id := range TransactionIDs {
		_, err := bufferedWriter.WriteString(id + "\n")
		if err != nil {
			log.Fatalf("Failed to write to file: %s", err)
		}
	}

	// Flush the buffered writer
	if err := bufferedWriter.Flush(); err != nil {
		log.Fatalf("Failed to flush buffered writer: %s", err)
	}

}
