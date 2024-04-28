package main

import (
	"crypto/sha256"
	"encoding/hex"
	"shramanpaul/utils"
)

func hash256(data string) string {
	// Decode hexadecimal string to byte slice
	rawBytes, _ := hex.DecodeString(data)

	// First SHA256 hash
	hash1 := sha256.Sum256(rawBytes)

	// Second SHA256 hash
	hash2 := sha256.Sum256(hash1[:])

	// Convert hash2 to hexadecimal string
	hashedString := hex.EncodeToString(hash2[:])

	return hashedString
}

func generateMerkleRoot(txids []string) string {
	if len(txids) == 0 {
		return ""
	}

	level := make([]string, len(txids))
	for idx := range txids {
		hash, _ := hex.DecodeString(txids[idx])
		hash = utils.ReverseBytes(hash)
		level[idx] = hex.EncodeToString(hash)
	}

	for len(level) > 1 {
		nextLevel := make([]string, 0)

		for i := 0; i < len(level); i += 2 {
			var pairHash string
			if i+1 == len(level) {
				pairHash = hash256(level[i] + level[i])
			} else {
				pairHash = hash256(level[i] + level[i+1])
			}
			nextLevel = append(nextLevel, pairHash)
		}

		level = nextLevel
	}

	return level[0]
}
