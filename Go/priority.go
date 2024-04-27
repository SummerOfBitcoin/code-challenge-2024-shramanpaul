package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func CalculateFee(tx Transaction) int {
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

func CalculateWeight(tx Transaction) int {
	baseSize := 0
	totalSize := 0

	// Add size for version and locktime
	baseSize += 4 + 4
	totalSize += 4 + 4

	// Add size for input count and output count
	inputCountSize := len(serialise_pubkey_len(uint64(len(tx.Vin))))
	outputCountSize := len(serialise_pubkey_len(uint64(len(tx.Vout))))
	baseSize += inputCountSize + outputCountSize
	totalSize += inputCountSize + outputCountSize

	// Calculate base size
	for _, input := range tx.Vin {
		baseSize += len(input.TxID)/2 + 34 // TxID size + output size (scriptpubkey + value)
	}
	for _, output := range tx.Vout {
		baseSize += 8 + 1 + len(output.Scriptpubkey)/2 // value + script length + scriptpubkey
	}

	// Calculate total size
	if IsSegWit(&tx) == 1 {
		totalSize += 2 // marker and flag
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

	// Calculate weight
	weight := baseSize*3 + totalSize
	return weight
}

// create an array of int for stroing the weight
var weight []int
var fees []int
var ratio []float64
var count, count2 int

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

		var tx Transaction

		err = json.Unmarshal(data, &tx)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err) // Print any errors
			continue
		}

		feeToWeightRatio := float64(CalculateFee(tx)) / float64(CalculateWeight(tx))
		if feeToWeightRatio >= 3.0 && CalculateWeight(tx) < 5200 {
			count++
			ratio = append(ratio, feeToWeightRatio)
			// fmt.Println("Ratio: ", feeToWeightRatio)
		}
	}
	// Convert the weight slice from []int to []string
	weightStrings := make([]string, len(ratio))
	for i, w := range ratio {
		weightStrings[i] = strconv.FormatFloat(w, 'f', -1, 64)
	}

	writeToFile(weightStrings)
	fmt.Println("count: ", count)
	// fmt.Println("count fees: ", count2)
}
