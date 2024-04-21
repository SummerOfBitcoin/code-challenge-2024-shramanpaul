package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Prevout3 struct {
	Value int `json:"value"`
}

type Input3 struct {
	Prevout Prevout3 `json:"prevout"`
}

type Output3 struct {
	Value int `json:"value"`
}

type Transaction3 struct {
	Inputs  []Input3  `json:"vin"`
	Outputs []Output3 `json:"vout"`
}

func Amount() int {
	files, err := ioutil.ReadDir("../mempool")
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

		var tx Transaction3
		var tx1 Transaction

		err = json.Unmarshal(data, &tx)
		if err != nil {
			log.Println("Error unmarshalling JSON:", err) // Print any errors
			continue
		}
		err = json.Unmarshal(data, &tx1)
		if err != nil {
			log.Println("Error unmarshalling JSON:", err) // Print any errors
			continue
		}

		if CalculateWeight(tx1) <= 1000 {

			for _, input := range tx.Inputs {
				totalInput += input.Prevout.Value
			}

			for _, output := range tx.Outputs {
				totalOutput += output.Value
			}
		}
	}

	// log.Println("Total input:", totalInput)
	// log.Println("Total output:", totalOutput)
	transactionfees := totalInput - totalOutput

	// log.Println("Total tranasction fees: ", transactionfees)

	return transactionfees
}
