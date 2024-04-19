package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var SegTransactionIDs []string
var SegwitMerkleRoot string

func ReaderN() {
	// SegTransactionIDs = append(SegTransactionIDs, SegwitMerkleRootCoinbase) //done		

	files, err := ioutil.ReadDir("../mempool")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filePath := "../mempool/" + file.Name()
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		var tx Transaction

		err = json.Unmarshal(data, &tx)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err) // Print any errors
			continue
		}

		serilised, _ := serializeTransaction(&tx)
		hash := to_sha(to_sha(serilised))
		// hash = reverseBytes(hash)

		SegTransactionIDs = append(SegTransactionIDs, hex.EncodeToString(hash))
	}

	SegwitMerkleRoot = generateMerkleRoot(SegTransactionIDs)

	fmt.Println("Computed Merkle Root Normal:", SegwitMerkleRoot)
	writeToFile(SegTransactionIDs)
}
