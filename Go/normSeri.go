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
	TxIDs = nil
	files, err := os.ReadDir("../mempool")
	if err != nil {
		log.Fatal(err)
	}
	count := 0

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
		feeToWeightRatio := float64(CalculateFee(tx)) / float64(CalculateWeight(tx))
		if feeToWeightRatio >= 3.0 && CalculateWeight(tx) < 5000 {
			count++
			serilised, _ := serializeTransaction(&tx)
			hash := reverseBytes(to_sha(to_sha(serilised)))

			TxIDs = append(TxIDs, hex.EncodeToString(hash))
		}
	}
	// fmt.Println("count: ",count)
	TxIDs = append([]string{NormalSerialiseCBTX}, TxIDs...)
	NormalMerkleRoot = generateMerkleRoot(TxIDs)
	// fmt.Println("OK: ", len(TxIDs))
	// fmt.Println("Computed Merkle Root Normal:", NormalMerkleRoot)
	// writeToFile(TxIDs)
}
