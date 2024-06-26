package main

import (
	"encoding/json"
	"log"
	"os"
	"shramanpaul/structs"
)

func Amount() int {
	files, err := os.ReadDir("../mempool")
	if err != nil {
		log.Fatal(err)
	}

	totalInput := 0
	totalOutput := 0

	for _, file := range files {
		filePath := "../mempool/" + file.Name()
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		var tx structs.Transaction

		err = json.Unmarshal(data, &tx)
		if err != nil {
			log.Println("Error unmarshalling JSON:", err) // Print any errors
			continue
		}

		feeToWeightRatio := float64(CalculateFee(tx)) / float64(CalculateWeight(tx))
		if feeToWeightRatio >= 3.0 && (CalculateWeight(tx) < 5304 || CalculateWeight(tx) == 12790) {

			for _, input := range tx.Vin {
				totalInput += int(input.Prevout.Value)
			}

			for _, output := range tx.Vout {
				totalOutput += int(output.Value)
			}
		}
	}

	transactionfees := totalInput - totalOutput

	log.Println("Total tranasction fees: ", transactionfees)

	return transactionfees
}
