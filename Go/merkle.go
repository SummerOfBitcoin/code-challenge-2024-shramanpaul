package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// func reverseString(hash string) string {
// 	hash = string(reverseBytes([]byte(hash)))
// 	return hash
// }

// func hash256(input string) string {
//     fmt.Println("txid: ", input)
//     h1 := sha256.Sum256([]byte(input))
//     fmt.Printf("h1: %x\n", h1)
//     h2 := sha256.Sum256(h1[:])
//     fmt.Printf("h2: %x\n", h2)

//     return hex.EncodeToString(h2[:])
// }

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
    for i, txid := range txids {
        txidBytes, _ := hex.DecodeString(txid)
        reversedBytes1 := reverseBytes1(txidBytes)
        level[i] = hex.EncodeToString(reversedBytes1)
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

func reverseBytes1(data []byte) []byte {
    for i := 0; i < len(data)/2; i++ {
        data[i], data[len(data)-1-i] = data[len(data)-1-i], data[i]
    }
    return data
}
func CreateMerkleTree() {

	var TransactionIDs []string

	files, err := os.ReadDir("../mempool")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	count := 0
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			txData, err := jsonData("../mempool/" + file.Name())
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			// Unmarshal the transaction data
			var tx Transaction
			err = json.Unmarshal([]byte(txData), &tx)
			if err != nil {
				panic(fmt.Sprintf("Error: %v", err))
				// continue
			}

			// Serialize the transaction
			serialized, err := serializeTransaction(&tx)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			count++
			fmt.Println("Count: ", count)

			// fmt.Printf("Serialized transaction: %x\n", serialized)
			hash := to_sha(to_sha(serialized))
			hash = reverseBytes(hash)

			TransactionIDs = append(TransactionIDs, hex.EncodeToString(hash))
			fmt.Printf("Transaction ID: %x\n", hash)

		}

	}

	// Call generateMerkleRoot function with txids from the file
	// merkleRoot := generateMerkleRoot(TransactionIDs)

	// // Print the computed Merkle root
	// fmt.Println("Computed Merkle Root:", merkleRoot)
	// fmt.Printf("Transaction IDs: %v\n", TransactionIDs)
	// writeToFile(TransactionIDs)
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
