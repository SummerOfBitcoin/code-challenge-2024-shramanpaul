package utils

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"shramanpaul/structs"
	
)

func To_sha(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// thia function will be used to derive the txids of ALL the transactions you are INCLUDING in the BLOCK --> include all these derived txids in the merkle root
func SerializeTransaction(tx *structs.Transaction) ([]byte, error) {
	var serialized []byte

	// Serialize version
	versionBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(versionBytes, tx.Version)
	serialized = append(serialized, versionBytes...)

	// Serialize vin count
	vinCount := uint64(len(tx.Vin))
	serialized = append(serialized, Serialise_pubkey_len(vinCount)...)

	// Serialize vin
	for _, vin := range tx.Vin {
		txidBytes, err := hex.DecodeString(vin.TxID)
		if err != nil {
			return nil, err
		}
		serialized = append(serialized, ReverseBytes(txidBytes)...)

		voutBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(voutBytes, vin.Vout)
		serialized = append(serialized, voutBytes...)

		// Serialize scriptSig length (empty for now)
		serialized_byte, err := hex.DecodeString(vin.Scriptsig)
		if err != nil {
			return nil, err
		}
		serialized = append(serialized, Serialise_pubkey_len(uint64(len(serialized_byte)))...)

		serialized = append(serialized, serialized_byte...)
		// Serialize sequence
		sequenceBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(sequenceBytes, vin.Sequence)
		serialized = append(serialized, sequenceBytes...)
	}

	// Serialize vout count
	voutCount := uint64(len(tx.Vout))
	serialized = append(serialized, Serialise_pubkey_len(voutCount)...)

	// Serialize vout
	for _, vout := range tx.Vout {
		valueBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(valueBytes, vout.Value)
		serialized = append(serialized, valueBytes...)

		// Serialize scriptPubKey length
		scriptPubKeyLen := uint64(len(vout.Scriptpubkey) / 2) // Divide by 2 to get byte length
		serialized = append(serialized, Serialise_pubkey_len(scriptPubKeyLen)...)

		// Serialize scriptPubKey
		scriptPubKeyBytes, err := hex.DecodeString(vout.Scriptpubkey)
		if err != nil {
			return nil, err
		}
		serialized = append(serialized, scriptPubKeyBytes...)
	}

	locktimeBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(locktimeBytes, tx.Locktime)
	serialized = append(serialized, locktimeBytes...)
	return serialized, nil
}

//thia function will be used to derive the wtxids of ALL the transactions you are INCLUDING in the BLOCK --> wtxids of legacy == txids of legacy --> include all these wtxids in the witness merkle.

func SerializeSegwit(tx *structs.Transaction) ([]byte, error) {
	if IsSegWit(tx) == 1 {
		var serialized []byte

		// Serialize version
		versionBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(versionBytes, tx.Version)
		serialized = append(serialized, versionBytes...)

		// Serialize marker
		serialized = append(serialized, 0x00)

		// Serialize flag
		serialized = append(serialized, 0x01)

		// Serialize vin count
		vinCount := uint64(len(tx.Vin))
		serialized = append(serialized, Serialise_pubkey_len(vinCount)...)

		// Serialize vin
		for _, vin := range tx.Vin {
			txidBytes, err := hex.DecodeString(vin.TxID)
			if err != nil {
				return nil, err
			}
			serialized = append(serialized, ReverseBytes(txidBytes)...)

			voutBytes := make([]byte, 4)
			binary.LittleEndian.PutUint32(voutBytes, vin.Vout)
			serialized = append(serialized, voutBytes...)

			// Serialize scriptSig length (empty for now)
			serialized_byte, err := hex.DecodeString(vin.Scriptsig)
			if err != nil {
				return nil, err
			}
			serialized = append(serialized, Serialise_pubkey_len(uint64(len(serialized_byte)))...)

			serialized = append(serialized, serialized_byte...)
			// Serialize sequence
			sequenceBytes := make([]byte, 4)
			binary.LittleEndian.PutUint32(sequenceBytes, vin.Sequence)
			serialized = append(serialized, sequenceBytes...)
		}

		// Serialize vout count
		voutCount := uint64(len(tx.Vout))
		serialized = append(serialized, Serialise_pubkey_len(voutCount)...)

		// Serialize vout
		for _, vout := range tx.Vout {
			valueBytes := make([]byte, 8)
			binary.LittleEndian.PutUint64(valueBytes, vout.Value)
			serialized = append(serialized, valueBytes...)

			// Serialize scriptPubKey length
			scriptPubKeyLen := uint64(len(vout.Scriptpubkey) / 2) // Divide by 2 to get byte length
			serialized = append(serialized, Serialise_pubkey_len(scriptPubKeyLen)...)

			// Serialize scriptPubKey
			scriptPubKeyBytes, err := hex.DecodeString(vout.Scriptpubkey)
			if err != nil {
				return nil, err
			}
			serialized = append(serialized, scriptPubKeyBytes...)
		}

		//witness
		for _, vin := range tx.Vin {
			witnessCount := uint64(len(vin.Witness))
			serialized = append(serialized, Serialise_pubkey_len(witnessCount)...)
			for _, witness := range vin.Witness {
				witnessBytes, err := hex.DecodeString(witness)
				if err != nil {
					return nil, err
				}
				serialized = append(serialized, Serialise_pubkey_len(uint64(len(witnessBytes)))...)
				serialized = append(serialized, witnessBytes...)
			}

		}

		locktimeBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(locktimeBytes, tx.Locktime)
		serialized = append(serialized, locktimeBytes...)
		return serialized, nil
	} else {
		byte, err := SerializeTransaction(tx)
		// fmt.Println("Transaction is not a Segwit Transaction")
		return byte, err
	}
	// return nil, nil
}

func ReverseBytes(data []byte) []byte {
	length := len(data)
	for i := 0; i < length/2; i++ {
		data[i], data[length-i-1] = data[length-i-1], data[i]
	}
	return data
}

func Serialise_pubkey_len(n uint64) []byte {
	if n < 0xfd {
		// If n < 0xfd, just return it as a single byte.
		return []byte{byte(n)}
	} else if n <= 0xffff {
		// If n <= 0xffff, return 0xfd followed by n as 2 bytes.
		b := make([]byte, 3)
		b[0] = 0xfd
		binary.LittleEndian.PutUint16(b[1:], uint16(n))
		return b
	} else if n <= 0xffffffff {
		// If n <= 0xffffffff, return 0xfe followed by n as 4 bytes.
		b := make([]byte, 5)
		b[0] = 0xfe
		binary.LittleEndian.PutUint32(b[1:], uint32(n))
		return b
	} else {
		// Otherwise, return 0xff followed by n as 8 bytes.
		b := make([]byte, 9)
		b[0] = 0xff
		binary.LittleEndian.PutUint64(b[1:], n)
		return b
	}
}
func uint16ToBytes(n uint16) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, n)
	return bytes
}

func uint32ToBytes(n uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, n)
	return bytes
}

func uint64ToBytes(n uint64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, n)
	return bytes
}

func Serialize() {
	files, err := os.ReadDir("../mempool")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	count := 0
	count1 := 0
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			txData, err := JsonData("../mempool/" + file.Name())
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			// Unmarshal the transaction data
			var tx structs.Transaction
			err = json.Unmarshal([]byte(txData), &tx)
			if err != nil {
				panic(fmt.Sprintf("Error: %v", err))
				// continue
			}

			// Serialize the transaction
			serialized, err := SerializeTransaction(&tx)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			// Print the serialized transaction
			// fmt.Printf("Serialized transaction: %x\n", serialized)
			hash := To_sha(To_sha(serialized))
			hash = ReverseBytes(hash)
			// fmt.Printf("Transaction ID: %x\n", hash)

			file_name := To_sha(hash)
			count++

			// fmt.Printf("File Name: %x File No: %d \n", file_name, count)
			if file.Name() != hex.EncodeToString(file_name)+".json" {
				count1++
				// fmt.Printf("Transaction ID: %x\n", hash)
				// fmt.Printf("Serialized transaction: %x\n", serialized)
				fmt.Printf("Actual File Name: %s Output file name: %x  File No: %d \n", file.Name(), file_name, count)
			}
		}
	}
	fmt.Println(count1)
}


