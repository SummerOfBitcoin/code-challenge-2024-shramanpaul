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
// var SegTransactionIDs []string
var SegTransactionIDsS []string
// var SegwitMerkleRoot string
var SegwitMerkleRootS string

func Reader() {
	// count := 0

	// var SegTransactionIDsonly []string
	// SegTransactionIDs = append(SegTransactionIDs, SegwitMerkleRootCoinbase)//done
	SegTransactionIDsS = append(SegTransactionIDsS, "0000000000000000000000000000000000000000000000000000000000000000")
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
		//appending the coinbase transaction to the SegTransactionIDs
		// if IsSegWit(tx) == 1 {
		serilisedS, _ := SerializeSegwit(&tx)
		// 	// fmt.Printf("Segwitserilisation: %x\n", serilised)
		hashS := to_sha(to_sha(serilisedS))
		// 	// hash = reverseBytes(hash)

		SegTransactionIDsS = append(SegTransactionIDsS, hex.EncodeToString(hashS))
		// 	SegTransactionIDsonly = append(SegTransactionIDsonly, hex.EncodeToString(hash))
		// fmt.Printf("Transaction ID: %x\n", hashS)
		// count++
		// fmt.Println("Count: ", count)
		// } else {
		// serilised, _ := serializeTransaction(&tx)
		// hash := to_sha(to_sha(serilised))
		// hash = reverseBytes(hash)

		// SegTransactionIDs = append(SegTransactionIDs, hex.EncodeToString(hash))
		// fmt.Printf("Transaction ID: %x\n", hash)
		// 	count++
		// fmt.Println("Count: ", count)

		// }
	}
	// SegwitMerkleRoot = generateMerkleRoot(SegTransactionIDs)
	SegwitMerkleRootS = generateMerkleRoot(SegTransactionIDsS)
	// SegwitonlyMerkleroot = generateMerkleRoot(SegTransactionIDsonly)

	// fmt.Println("Segwit Merkle Root:", SegwitonlyMerkleroot)

	// fmt.Println("Computed Merkle Root:", SegwitMerkleRoot)

	// fmt.Println("Computed Merkle RootS:", SegwitMerkleRootS)

	// writeToFile(SegTransactionIDsonly)
	// writeToFile(SegTransactionIDs)
	// writeToFile(SegTransactionIDsS)
	// fmt.Println("Segwit Transaction IDs: ", SegTransactionIDs)
}
