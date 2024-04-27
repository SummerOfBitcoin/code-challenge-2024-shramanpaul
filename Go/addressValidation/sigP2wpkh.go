package addressValidation

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"shramanpaul/structs"
	"shramanpaul/utils"
	"strings"

	"golang.org/x/crypto/ripemd160"
)

func SigVerify2() {
	files, err := ioutil.ReadDir("../mempool/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join("../mempool", file.Name())
			data, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			var tx structs.Transaction
			err = json.Unmarshal(data, &tx)
			if err != nil {
				fmt.Println("Error unmarshalling JSON:", err)
				continue
			}

			for _, vin := range tx.Vin {
				if vin.Prevout.ScriptpubkeyType == "v0_p2wpkh" {
					fmt.Println(filePath)
					SerializeP2wpkh(filePath)
					break
				} else {
					break
				}
			}
		}
	}
}

func SerializeP2wpkh(filename string) {
	txData, err := utils.JsonData(filename)
	Handle(err)

	// Unmarshal the transaction data
	var tx structs.Transaction
	err = json.Unmarshal([]byte(txData), &tx)
	if err != nil {
		panic(fmt.Sprintf("Error: %v", err))
	}

	// for idx := range tx.Vin {
	version := tx.Version
	fmt.Println("Version: ", version)

	// Initialize an empty string to hold the serialized data
	var serializedData string

	// Iterate over the inputs
	for _, input := range tx.Vin {
		// Convert Vout to a byte slice
		voutBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(voutBytes, uint32(input.Vout))

		// Convert the byte slice to a hexadecimal string
		voutHex := hex.EncodeToString(voutBytes)

		// Serialize the TXID and VOUT
		serializedData += input.TxID + voutHex
	}

	fmt.Println("Serialized data:", serializedData)

	// Hash the serialized data using SHA-256
	ser, _ := hex.DecodeString(serializedData)
	hash := sha256.Sum256(ser)
	hash = sha256.Sum256(hash[:])

	// Print the hashed data
	fmt.Println("Hashed data:", hex.EncodeToString(hash[:]))

	// Initialize an empty string to hold the serialized data
	var serializedDataSeq string

	// Iterate over the inputs
	for _, input := range tx.Vin {
		// Convert Sequence to a byte slice
		seqBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(seqBytes, uint32(input.Sequence))

		// Convert the byte slice to a hexadecimal string
		seqHex := hex.EncodeToString(seqBytes)

		// Serialize the sequence
		serializedDataSeq += seqHex
	}

	fmt.Println("Serialized sequence data:", serializedDataSeq)

	// Hash the serialized data using SHA-256
	serSeq, _ := hex.DecodeString(serializedDataSeq)
	hashSeq := sha256.Sum256(serSeq)
	hashSeq = sha256.Sum256(hashSeq[:])

	// Print the hashed data
	fmt.Println("Hashed data:", hex.EncodeToString(hashSeq[:]))

	// Extract the signature from the witness field in the Vin
	signature := tx.Vin[0].Witness[0]
	fmt.Println("Signature: ", signature)

	// Get the input you want to spend
	input := tx.Vin[0]

	// Extract the public key hash from the Scriptpubkey
	pubKeyHash := input.Prevout.Scriptpubkey[4:44] // Adjust these indices as needed
	fmt.Println("Public Key Hash: ", pubKeyHash)

	// Iterate over the vout array
	for _, output := range tx.Vout {
		// Check if the scriptpubkey contains the public key hash
		if strings.Contains(output.Scriptpubkey, pubKeyHash) {
			// Output the type of the vout
			// fmt.Println("Vout type: ", output.ScriptpubkeyType)
			// break
			if output.ScriptpubkeyType == "p2pkh" {
				fmt.Println(output.ScriptpubkeyType)
				// Create the scriptcode
				scriptCode := "1976a914" + pubKeyHash + "88ac"
				fmt.Println("Scriptcode: ", scriptCode)
			} else if output.ScriptpubkeyType == "v0_p2wpkh" {
				fmt.Println(output.ScriptpubkeyType)
				// Create the scriptcode
				scriptCode := "0014" + pubKeyHash
				fmt.Println("Scriptcode: ", scriptCode)
			}
		}
	}

	//take the input amount
	amount := tx.Vin[0].Prevout.Value
	fmt.Println("Amount: ", amount)

	// Iterate over the transaction inputs
	for i, input := range tx.Vin {
		if len(input.Witness) >= 2 {
			// Extract the signature and public key from the witness field in the Vin
			witness := input.Witness
			signature := witness[len(witness)-2]
			publicKey := witness[len(witness)-1]

			// Construct the witness field
			witnessField := "02" + signature + "01" + publicKey
			fmt.Printf("Witness field for input %d: %s\n", i, witnessField)
		} else {
			fmt.Printf("Witness field for input %d has less than 2 elements\n", i)
		}
	}

	// Serialize the transaction
	serialized, err := utils.SerializeTransaction(&tx)
	if err != nil {
		log.Fatalf("Failed to serialize transaction: %v", err)
	}

	// Append the SIGHASH_ALL flag to the serialized transaction
	sighash := make([]byte, 4)
	binary.LittleEndian.PutUint32(sighash, 1)
	serialized = append(serialized, sighash...)

	fmt.Printf("Serialized transaction: %x\n", serialized)

	// Double SHA-256 hash the serialized transaction
	hash2 := sha256.Sum256(serialized)
	hash2 = sha256.Sum256(hash2[:])

	fmt.Printf("Double SHA-256 hash of serialized transaction: %x\n", hash2)

	// Extract the public key from the witness field
	publicKey := tx.Vin[0].Witness[len(tx.Vin[0].Witness)-1]

	// Decode the public key
	pubKeyBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		fmt.Printf("Failed to decode public key: %v\n", err)
		return
	}

	// Compute the hash160 of the public key
	hasher := sha256.New()
	hasher.Write(pubKeyBytes)
	hashResult := hasher.Sum(nil)
	hasher = ripemd160.New()
	hasher.Write(hashResult)
	pubKeyHashResult := hasher.Sum(nil)
	pubKeyHash = hex.EncodeToString(pubKeyHashResult)

	// Check if the hash160 of the public key matches the hash in the scriptpubkey
	scriptPubKey := tx.Vin[0].Prevout.Scriptpubkey[4:]
	if pubKeyHash == scriptPubKey {
		fmt.Println("The witness field is correct.")
	} else {
		fmt.Println("The witness field is incorrect.")
		return
	}
}
