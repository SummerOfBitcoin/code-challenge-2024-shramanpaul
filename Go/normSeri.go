package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var TxIDs []string
var NormalMerkleRoot string

func ReaderN() {
	// TxIDs = append(TxIDs, NormalSerialiseCBTX) //done

	files, err := os.ReadDir("../mempool")
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
		if CalculateWeight(tx) <= 720 {

			serilised, _ := serializeTransaction(&tx)
			hash := reverseBytes(to_sha(to_sha(serilised)))

			TxIDs = append(TxIDs, hex.EncodeToString(hash))
		}
	}
	TxIDs = append([]string{NormalSerialiseCBTX}, TxIDs...)
	NormalMerkleRoot = generateMerkleRoot(TxIDs)
	fmt.Println("OK: ", len(TxIDs))
	fmt.Println("Computed Merkle Root Normal:", NormalMerkleRoot)
	writeToFile(TxIDs)
}
