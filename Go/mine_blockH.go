package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"shramanpaul/structs"
	"time"
)
//Block Header and Mining
//block header is made in this and then mining is also done by incrementing nonce

var BlockHeaderHash string
var BlockHeaderHex string

func PrintBlockHeader() {
	version := int32(0x20000000) 								//version
	var previousBlock [32]byte  								//previous block hash 
	var merkleRoot [32]byte		 								//merkle root
	merkleRootHex := NormalMerkleRoot
	fmt.Println("Normal MerkleRoot: ", NormalMerkleRoot)
	merkleRootBytes, _ := hex.DecodeString(merkleRootHex)
	copy(merkleRoot[:], (merkleRootBytes))

	timestamp := uint32(time.Now().Unix()) 						//timestamp

	bitsHex := "1f00ffff"										//difficulty target in compact form
	bitsBytes, _ := hex.DecodeString(bitsHex)
	reverseBytes(bitsBytes)
	bits := binary.BigEndian.Uint32(bitsBytes)

	nonce := uint32(0)											//nonce

	blockHeader := &structs.BlockHeader{
		Version:       version,
		PreviousBlock: previousBlock,
		MerkleRoot:    merkleRoot,
		Timestamp:     timestamp,
		Bits:          bits,
		Nonce:         nonce,
	}

	versionBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(versionBytes, uint32(blockHeader.Version))

	timestampBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(timestampBytes, blockHeader.Timestamp)

	bitBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bitBytes, blockHeader.Bits)

	nonceBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(nonceBytes, blockHeader.Nonce)

	blockHeaderBytes := append(versionBytes, blockHeader.PreviousBlock[:]...)
	blockHeaderBytes = append(blockHeaderBytes, blockHeader.MerkleRoot[:]...)
	blockHeaderBytes = append(blockHeaderBytes, timestampBytes...)
	blockHeaderBytes = append(blockHeaderBytes, bitsBytes...)
	blockHeaderBytes = append(blockHeaderBytes, nonceBytes...)

	hash := to_sha(to_sha(blockHeaderBytes))
	hashInt := new(big.Int).SetBytes(hash)

	// Check the block header hash against the difficulty target
	difficultyTarget := "0000ffff00000000000000000000000000000000000000000000000000000000"
	difficultyTargetBytes, _ := hex.DecodeString(difficultyTarget)
	difficultyTargetInt := new(big.Int).SetBytes((difficultyTargetBytes))

	blockHeader.Nonce = 0 
	for {
		nonceBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(nonceBytes, blockHeader.Nonce)
		blockHeaderBytes = append(blockHeaderBytes[:76], nonceBytes...) // Update the nonce in the block header bytes

		hash = to_sha(to_sha(blockHeaderBytes))
		hash = reverseBytes(hash)
		hashInt.SetBytes((hash))

		if hashInt.Cmp(difficultyTargetInt) <= 0 {
			// The block header hash is less than or equal to the difficulty target, so the nonce is valid
			break
		}
		blockHeader.Nonce++
	}
	BlockHeaderHex = hex.EncodeToString(blockHeaderBytes)
	fmt.Println("BlockHeader: ", BlockHeaderHex)
	hash = reverseBytes(hash)
	BlockHeaderHash = hex.EncodeToString(hash)

}

func Block() {
	PrintBlockHeader()
}
