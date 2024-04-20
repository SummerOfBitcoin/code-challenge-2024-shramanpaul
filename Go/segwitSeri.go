package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// var SegwitonlyMerkleroot string
// var TxIDs []string
var WtxIDs []string

// var SegwitMerkleRoot string
var SegwitMerkleRootS string

func Reader() {
	// count := 0

	// var SegTransactionIDsonly []string
	// TxIDs = append(TxIDs, NormalSerialiseCBTX)//done
	WtxIDs = append(WtxIDs, "0000000000000000000000000000000000000000000000000000000000000000")
	// SegTransactionIDsonly = append(SegTransactionIDsonly, "0000000000000000000000000000000000000000000000000000000000000000")

	files, err := ioutil.ReadDir("../mempool")
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
		//appending the coinbase transaction to the TxIDs
		// if IsSegWit(tx) == 1 {
		serilisedS, _ := SerializeSegwit(&tx)
		// 	// fmt.Printf("Segwitserilisation: %x\n", serilised)
		hashS := to_sha(to_sha(serilisedS))
		// 	// hash = reverseBytes(hash)

		WtxIDs = append(WtxIDs, hex.EncodeToString(hashS))
		// 	SegTransactionIDsonly = append(SegTransactionIDsonly, hex.EncodeToString(hash))
		// fmt.Printf("Transaction ID: %x\n", hashS)
		// count++
		// fmt.Println("Count: ", count)
		// } else {
		// serilised, _ := serializeTransaction(&tx)
		// hash := to_sha(to_sha(serilised))
		// hash = reverseBytes(hash)

		// TxIDs = append(TxIDs, hex.EncodeToString(hash))
		// fmt.Printf("Transaction ID: %x\n", hash)
		// 	count++
		// fmt.Println("Count: ", count)

		// }
	}
	// SegwitMerkleRoot = generateMerkleRoot(TxIDs)
	commitmentHeader := "6a24aa21a9ed"
	SegwitMerkleRootS = generateMerkleRoot(WtxIDs)
	WitnessReservedValue := "0000000000000000000000000000000000000000000000000000000000000000"
	// Decode the hexadecimal strings to bytes
	commitmentHeaderH, _ := hex.DecodeString(commitmentHeader)
	SegwitMerkleRootH, _ := hex.DecodeString(SegwitMerkleRootS)
	WitnessReserved, _ := hex.DecodeString(WitnessReservedValue)

	// Concatenate and hash the bytes
	hash := to_sha(to_sha(append(SegwitMerkleRootH, WitnessReserved...)))
	hash = append(commitmentHeaderH, hash...)

	// Encode the hash to a hexadecimal string
	SegwitMerkleRootS = hex.EncodeToString(hash)
	fmt.Println("ScriptPubkey of CBTX:", SegwitMerkleRootS)
	// SegwitonlyMerkleroot = generateMerkleRoot(SegTransactionIDsonly)

	// fmt.Println("Segwit Merkle Root:", SegwitonlyMerkleroot)

	// fmt.Println("Computed Merkle Root:", SegwitMerkleRoot)

	// fmt.Println("Computed Merkle RootS:", SegwitMerkleRootS)

	// writeToFile(SegTransactionIDsonly)
	// writeToFile(TxIDs)
	// writeToFile(WtxIDs)
	// fmt.Println("Segwit Transaction IDs: ", TxIDs)
}
