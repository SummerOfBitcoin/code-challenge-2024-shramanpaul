package main

import (
	"log"
	"os"
	"strings"

	"github.com/mr-tron/base58"
)

func to_byte(data string) []byte {
	return []byte(data)
}
func Base58Encode(input []byte) []byte {
	var data string = base58.Encode(input)
	return []byte(data)
}
func Handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// addressValidation.Address_Validation()
	// Serialize()
	// SigVerify()
	// SigVerify2()
	// SaveTxids()
	// CreateMerkleTree()
	// Segwit()
	Reader()
	Cointransaction()
	// Amount()
	ReaderN()
	Block()

	// Open the file in append mode, or create it if it doesn't exist
	file, err := os.OpenFile("../output.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()

	// Write BlockHeaderHash to the file
	_, err = file.WriteString(BlockHeaderHex + "\n")
	if err != nil {
		log.Fatalf("Failed writing to file: %s", err)
	}

	// Write SerialisedCBTX to the file
	_, err = file.WriteString(SerialisedCBTX + "\n")
	if err != nil {
		log.Fatalf("Failed writing to file: %s", err)
	}

	// Write TXID to the file
	//reversebyte SegTransactionIDs
	for i := 0; i < len(SegTransactionIDs); i++ {
		SegTransactionIDs[i] = string(reverseBytes([]byte(SegTransactionIDs[i])))
	}

	_, err = file.WriteString(strings.Join(SegTransactionIDs, "\n") + "\n")
	if err != nil {
		log.Fatalf("Failed writing to file: %s", err)
	}

}
