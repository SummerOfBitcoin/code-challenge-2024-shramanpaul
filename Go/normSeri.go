package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"shramanpaul/structs"
	"shramanpaul/utils"
)

var TxIDs []string
var NormalMerkleRoot string

func ReaderN() {
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

		var tx structs.Transaction

		err = json.Unmarshal(data, &tx)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err) // Print any errors
			continue
		}
		feeToWeightRatio := float64(CalculateFee(tx)) / float64(CalculateWeight(tx))
		if feeToWeightRatio >= 3.0 && (CalculateWeight(tx) < 5304 || CalculateWeight(tx) == 12790) {
			count++
			serilised, _ := utils.SerializeTransaction(&tx)
			hash := utils.ReverseBytes(utils.To_sha(utils.To_sha(serilised)))

			TxIDs = append(TxIDs, hex.EncodeToString(hash))
		}
	}

	TxIDs = append([]string{NormalSerialiseCBTX}, TxIDs...)
	NormalMerkleRoot = generateMerkleRoot(TxIDs)

}
