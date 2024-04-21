package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
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

		var tx Transaction

		err = json.Unmarshal(data, &tx)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err) // Print any errors
			continue
		}
		if CalculateWeight(tx) <= 4000 {
			count++
			serilisedS, _ := SerializeSegwit(&tx)
			// fmt.Println("Segwit serilisedS: ",hex.EncodeToString(serilisedS))
			// fmt.Println("count of tx:",count)
			hashS := reverseBytes(to_sha(to_sha(serilisedS)))

			WtxIDs = append(WtxIDs, hex.EncodeToString(hashS))
		}
	}
	// fmt.Println("count: ",count)
	commitmentHeader := "6a24aa21a9ed"
	WtxIDs = append([]string{"0000000000000000000000000000000000000000000000000000000000000000"}, WtxIDs...)
	SegwitMerkleRootS = generateMerkleRoot(WtxIDs)
	// fmt.Println("OK segwit: ",len(WtxIDs))
	// fmt.Println("SegwitMerkleRoot: ",SegwitMerkleRootS)

	WitnessReservedValue := "0000000000000000000000000000000000000000000000000000000000000000"

	// Decode the hexadecimal strings to bytes
	commitmentHeaderH, _ := hex.DecodeString(commitmentHeader)
	SegwitMerkleRootH, _ := hex.DecodeString(SegwitMerkleRootS)
	WitnessReserved, _ := hex.DecodeString(WitnessReservedValue)

	// Concatenate and hash the bytes
	fmt.Println("Witness root hash: ", hex.EncodeToString(SegwitMerkleRootH))
	// fmt.Println("Segwit root: ", hex.EncodeToString(append(SegwitMerkleRootH, WitnessReserved...)))
	hash := to_sha(to_sha(append(SegwitMerkleRootH, WitnessReserved...)))
	hash = append(commitmentHeaderH, hash...)

	SegwitMerkleRootS = hex.EncodeToString(hash)
	fmt.Println("Computed Merkle Root Segwit: ", SegwitMerkleRootS)

	writeToFile(WtxIDs)

}
