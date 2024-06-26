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

var WtxIDs []string

var SegwitMerkleRootS string

func Reader() {
	WtxIDs = nil
	// WtxIDs = append(WtxIDs, "00000000000000000000000000000000000000000000000000000000000000")
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
			serilisedS, _ := utils.SerializeSegwit(&tx)
	
			hashS := utils.ReverseBytes(utils.To_sha(utils.To_sha(serilisedS)))

			WtxIDs = append(WtxIDs, hex.EncodeToString(hashS))
		}
	}
	// fmt.Println("count: ",count)
	commitmentHeader := "6a24aa21a9ed"
	WtxIDs = append([]string{"0000000000000000000000000000000000000000000000000000000000000000"}, WtxIDs...)
	SegwitMerkleRootS = generateMerkleRoot(WtxIDs)


	WitnessReservedValue := "0000000000000000000000000000000000000000000000000000000000000000"

	commitmentHeaderH, _ := hex.DecodeString(commitmentHeader)
	SegwitMerkleRootH, _ := hex.DecodeString(SegwitMerkleRootS)
	WitnessReserved, _ := hex.DecodeString(WitnessReservedValue)


	fmt.Println("Witness root hash: ", hex.EncodeToString(SegwitMerkleRootH))
	
	hash := utils.To_sha(utils.To_sha(append(SegwitMerkleRootH, WitnessReserved...)))
	hash = append(commitmentHeaderH, hash...)

	SegwitMerkleRootS = hex.EncodeToString(hash)
	fmt.Println("Computed Merkle Root Segwit: ", SegwitMerkleRootS)

	

}
