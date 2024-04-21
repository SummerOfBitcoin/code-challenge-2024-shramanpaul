package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type PriorityTransaction struct {
	ID     string
	Fee    int
	Weight int
}

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

	// Calculate base size
	for _, input := range tx.Vin {
		baseSize += len(input.TxID)/2 + 34 // TxID size + output size (scriptpubkey + value)
		baseSize += 4 + 1 + 4              // version + input count + locktime
	}
	for _, output := range tx.Vout {
		baseSize += 8 + 1 + len(output.Scriptpubkey)/2 // value + script length + scriptpubkey
		baseSize += 1 + 4                              // output count + locktime
	}

	// Calculate total size
	totalSize += 2 // marker and flag
	for _, input := range tx.Vin {
		totalSize += len(input.TxID)/2 + 34 // TxID size + output size (scriptpubkey + value)
		totalSize += 4 + 1 + 4              // version + input count + locktime
		for _, witness := range input.Witness {
			totalSize += len(witness) / 2 // witness size
		}
	}
	for _, output := range tx.Vout {
		totalSize += 8 + 1 + len(output.Scriptpubkey)/2 // value + script length + scriptpubkey
		totalSize += 1 + 4                              // output count + locktime
	}

	// Calculate weight
	weight := baseSize*3 + totalSize

	return weight
}

// create an array of int for stroing the weight
var weight []int
var count int

func Priority() {
	// files, err := os.ReadDir("../mempool")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// priorityTransactions := []PriorityTransaction{}

	// for _, file := range files {
	// 	filePath := "../mempool/" + file.Name()
	// 	data, err := os.ReadFile(filePath)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	var tx Transaction

	// 	err = json.Unmarshal(data, &tx)
	// 	if err != nil {
	// 		fmt.Println("Error unmarshalling JSON:", err) // Print any errors
	// 		continue
	// 	}

	// 	priorityTx := PriorityTransaction{
	// 		ID:     tx.ID,
	// 		Fee:    CalculateFee(tx),
	// 		Weight: CalculateWeight(tx),
	// 	}
	// 	priorityTransactions = append(priorityTransactions, priorityTx)
	// }

	// // Sort the transactions by fee per weight unit in descending order
	// sort.Slice(priorityTransactions, func(i, j int) bool {
	// 	fpwuI := float64(priorityTransactions[i].Fee) / float64(priorityTransactions[i].Weight)
	// 	fpwuJ := float64(priorityTransactions[j].Fee) / float64(priorityTransactions[j].Weight)
	// 	return fpwuI > fpwuJ
	// })

	// // Print the sorted transactions
	// for _, tx := range priorityTransactions {
	// 	fmt.Printf("ID: %s, Fee: %d, Weight: %d\n", tx.ID, tx.Fee, tx.Weight)
	// }

	//read a specific file
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
		if CalculateWeight(tx) <= 600 {
			count++

			weight = append(weight, CalculateWeight(tx))
		}
		// fmt.Println("Fee:", CalculateFee(tx))
		// fmt.Println("Weight:", (CalculateWeight(tx)))
	}
	fmt.Println("count: ", count)
	// Convert the weight slice to a slice of strings
	var weightStrings []string
	for _, w := range weight {
		weightStrings = append(weightStrings, strconv.Itoa(w))
	}

	// writeToFile(weightStrings)
}
