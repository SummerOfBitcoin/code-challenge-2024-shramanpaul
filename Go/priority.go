package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"shramanpaul/structs"
	"shramanpaul/utils"
	"strconv"
)

func CalculateFee(tx structs.Transaction) int {
	totalInput := 0
	totalOutput := 0

	for _, input := range tx.Vin {
		totalInput += int(input.Prevout.Value)
	}

	for _, output := range tx.Vout {
		totalOutput += int(output.Value)
	}
	return totalInput - totalOutput
}

func CalculateWeight(tx structs.Transaction) int {
	baseSize := 0
	totalSize := 0

	
	baseSize += 4 + 4
	totalSize += 4 + 4

	
	inputCountSize := len(utils.Serialise_pubkey_len(uint64(len(tx.Vin))))
	outputCountSize := len(utils.Serialise_pubkey_len(uint64(len(tx.Vout))))
	baseSize += inputCountSize + outputCountSize
	totalSize += inputCountSize + outputCountSize

	
	for _, input := range tx.Vin {
		baseSize += len(input.TxID)/2 + 34 // TxID size + output size (scriptpubkey + value)
	}
	for _, output := range tx.Vout {
		baseSize += 8 + 1 + len(output.Scriptpubkey)/2 // value + script length + scriptpubkey
	}


	if utils.IsSegWit(&tx) == 1 {
		totalSize += 2 
		for _, input := range tx.Vin {
			totalSize += len(input.TxID)/2 + 34 // TxID size + output size (scriptpubkey + value)
			for _, witness := range input.Witness {
				totalSize += len(witness) / 2 // witness size
			}
		}
		for _, output := range tx.Vout {
			totalSize += 8 + 1 + len(output.Scriptpubkey)/2 // value + script length + scriptpubkey
		}
	} else {
		totalSize = baseSize
	}

	weight := baseSize*3 + totalSize
	return weight
}


var ratio []float64
var count int

func Priority() {

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

		var tx structs.Transaction

		err = json.Unmarshal(data, &tx)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err) // Print any errors
			continue
		}

		feeToWeightRatio := float64(CalculateFee(tx)) / float64(CalculateWeight(tx))
		if feeToWeightRatio >= 3.0 && (CalculateWeight(tx) < 5304 || CalculateWeight(tx) == 12790) {
			count++
			ratio = append(ratio, feeToWeightRatio)
			
		}
	}
	
	weightStrings := make([]string, len(ratio))
	for i, w := range ratio {
		weightStrings[i] = strconv.FormatFloat(w, 'f', -1, 64)
	}

	fmt.Println("count: ", count)
	
}
