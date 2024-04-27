package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"shramanpaul/utils"
)

func hash256(data string) string {
	// Decode hexadecimal string to byte slice
	rawBytes, _ := hex.DecodeString(data)

	// First SHA256 hash
	hash1 := sha256.Sum256(rawBytes)

	// Second SHA256 hash
	hash2 := sha256.Sum256(hash1[:])

	// Convert hash2 to hexadecimal string
	hashedString := hex.EncodeToString(hash2[:])

	return hashedString
}

func generateMerkleRoot(txids []string) string {
	if len(txids) == 0 {
		return ""
	}

	level := make([]string, len(txids))
	for idx := range txids {
		hash, _ := hex.DecodeString(txids[idx])
		hash = utils.ReverseBytes(hash)
		level[idx] = hex.EncodeToString(hash)
	}

	for len(level) > 1 {
		nextLevel := make([]string, 0)

		for i := 0; i < len(level); i += 2 {
			var pairHash string
			if i+1 == len(level) {
				pairHash = hash256(level[i] + level[i])
			} else {
				pairHash = hash256(level[i] + level[i+1])
			}
			nextLevel = append(nextLevel, pairHash)
		}

		level = nextLevel
	}

	return level[0]
}

// Write the TransactionIDs to a file
func writeToFile(TransactionIDs []string) {
	// Open the file in write-only mode. If the file doesn't exist, create it.
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()

	// Create a buffered writer from the file
	bufferedWriter := bufio.NewWriter(file)

	// Write the TransactionIDs to the file
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
