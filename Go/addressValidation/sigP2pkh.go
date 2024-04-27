// package main

// import (
// 	"encoding/asn1"
// 	"encoding/binary"
// 	"encoding/hex"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"math/big"
// 	"strconv"
// )

// func extractSubstring(s string, startIndex, endIndex int) string {
// 	if startIndex < 0 || endIndex > len(s) {
// 		return "" // Return an empty string if the indices are out of bounds
// 	}
// 	return s[startIndex:endIndex]
// }

// type ecdsaSignature struct {
// 	R, S *big.Int
// }

// func extractRandS(sig string) (*big.Int, *big.Int, error) {
// 	sigBytes, err := hex.DecodeString(sig)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	var ecdsaSig ecdsaSignature
// 	_, err = asn1.Unmarshal(sigBytes, &ecdsaSig)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return ecdsaSig.R, ecdsaSig.S, nil
// }

// func SigVerify() {
// 	files, err := ioutil.ReadDir("../mempool/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, file := range files {
// 		if !file.IsDir() {
// 			filePath := "../mempool/" + file.Name()
// 			data, err := ioutil.ReadFile(filePath)
// 			if err != nil {
// 				fmt.Println("Error reading file:", err) // Print any errors
// 				continue
// 			}

// 			var tx Transaction
// 			err = json.Unmarshal(data, &tx)
// 			if err != nil {
// 				fmt.Println("Error unmarshalling JSON:", err) // Print any errors
// 				continue
// 			}

// 			for idx := range tx.Vin {
// 				if tx.Vin[idx].Prevout.ScriptpubkeyType == "p2pkh" {
// 					fmt.Println(filePath)
// 					SerializeP2ph(filePath)
// 				}
// 			}
// 		}
// 	}
// }
// func SerializeP2ph(filename string) {

// 	count := 0 // Move the count declaration outside the loop

// 	txData, err := jsonData(filename)
// 	Handle(err)

// 	// Unmarshal the transaction data
// 	var tx Transaction
// 	err = json.Unmarshal([]byte(txData), &tx)
// 	if err != nil {
// 		panic(fmt.Sprintf("Error: %v", err))
// 		// continue
// 	}
// 	var scriptsigcopy string
// 	// Serialize the transaction
// 	for idx := range tx.Vin {
// 		fmt.Println("ScriptPubKeyType: ", tx.Vin[idx].Prevout.ScriptpubkeyType)
// 		if tx.Vin[idx].Prevout.ScriptpubkeyType == "p2pkh" || tx.Vin[idx].Prevout.ScriptpubkeyType == "p2sh" {

// 			scriptsigcopy = tx.Vin[idx].Scriptsig
// 			fmt.Printf("ScritpSig: %s\n", scriptsigcopy) //copy of scriptsig

// 			// Extract the first two digits from the substring
// 			firstTwoDigits := scriptsigcopy[:2]

// 			// Convert the first two digits to an integer
// 			firstTwoDigitsInt, err := strconv.Atoi(firstTwoDigits)
// 			if err != nil {
// 				fmt.Println("Error:", err)
// 				return
// 			}
// 			finalpos := 0
// 			if firstTwoDigitsInt%2 == 1 {
// 				finalpos = firstTwoDigitsInt*3 + 1
// 			} else {
// 				finalpos = firstTwoDigitsInt * 3
// 			}
// 			fmt.Println("First two digits: ", firstTwoDigitsInt)

// 			substring := extractSubstring(scriptsigcopy, 2, finalpos)
// 			fmt.Println("Signature: ", substring) //correct signature

// 			sigAppendsighash := substring + "01"
// 			fmt.Println("Signature with sighash appended: ", sigAppendsighash)

// 			r, s, err := extractRandS(sigAppendsighash)
// 			if err != nil {
// 				fmt.Println("Error:", err)
// 				return
// 			}
// 			fmt.Println("R: ", r)
// 			fmt.Println("S: ", s)
// 			// The public key is the last 66 characters of the scriptsig_asm field
// 			publicKey := tx.Vin[idx].ScriptsigAsm[len(tx.Vin[idx].ScriptsigAsm)-66:]

// 			fmt.Println("Public Key: ", publicKey)
// 			// Convert sigAppendsighash and publicKey to byte slices
// 			sigBytes, err := hex.DecodeString(sigAppendsighash)
// 			if err != nil {
// 				fmt.Println("Error:", err)
// 				return
// 			}
// 			fmt.Printf("sigbytes: %x\n", sigBytes) //sigbytes
// 			pubkeyBytes, err := hex.DecodeString(publicKey)
// 			if err != nil {
// 				fmt.Println("Error:", err)
// 				return
// 			}

// 			// Prepend the length of sigAppendsighash and publicKey
// 			sigBytes = append([]byte{byte(len(sigBytes))}, sigBytes...)
// 			pubkeyBytes = append([]byte{byte(len(pubkeyBytes))}, pubkeyBytes...)

// 			// Append the byte slices together to form the scriptSig
// 			scriptSig := append(sigBytes, pubkeyBytes...)

// 			fmt.Printf("ScriptSig2: %x\n", scriptSig) //scriptsig2
// 			if hex.EncodeToString(scriptSig) != tx.Vin[idx].Scriptsig {
// 				fmt.Println("ScriptSig does not match")
// 				break
// 			}

// 			tx.Vin[idx].Scriptsig = tx.Vin[idx].Prevout.Scriptpubkey
// 		}
// 	}
// 	serialized, err := serializeTransaction(&tx)
// 	Handle(err)
// 	sighash := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(sighash, 1)
// 	serialized = append(serialized, sighash...)
// 	fmt.Printf("Serialized transaction: %x\n", serialized)
// 	serialized = to_sha(to_sha(serialized))
// 	fmt.Printf("Serialized d2 hash: %x\n", serialized)
// 	// Print the final count after the loop
// 	fmt.Printf("Total count: %d\n", count)
// }