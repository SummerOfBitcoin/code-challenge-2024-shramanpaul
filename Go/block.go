package main

import (
	"encoding/binary"
	"encoding/hex"
	"math/big"
	"math/rand"
	"time"
)

type BlockHeader struct {
	Version       int32
	PreviousBlock [32]byte
	MerkleRoot    [32]byte
	Timestamp     uint32
	Bits          uint32
	Nonce         uint32
}

var BlockHeaderHash string
var BlockHeaderHex string

func PrintBlockHeader() {
	version := int32(0x20000000) // Static version value
	var previousBlock [32]byte   // Empty byte array
	var merkleRoot [32]byte
	merkleRootHex := SegwitMerkleRoot
	merkleRootBytes, _ := hex.DecodeString(merkleRootHex)
	copy(merkleRoot[:], reverseBytes(merkleRootBytes))

	timestamp := uint32(time.Now().Unix()) //timestamp

	// difficultytarget:="0000ffff00000000000000000000000000000000000000000000000000000000"
	bitsHex := "1f00ffff"
	bitsBytes, _ := hex.DecodeString(bitsHex)
	reverseBytes(bitsBytes)
	bits := binary.BigEndian.Uint32(bitsBytes)

	nonce := rand.Uint32()

	blockHeader := &BlockHeader{
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

	// fmt.Println("BlockHeaderBytes: ",blockHeaderBytes)
	//convert blockheader to hex
	// blockHeaderHex := hex.EncodeToString(blockHeaderBytes)

	hash := to_sha(to_sha(blockHeaderBytes))
	hashInt := new(big.Int).SetBytes(hash)

	// Check the block header hash against the difficulty target
	difficultyTarget := "0000ffff00000000000000000000000000000000000000000000000000000000"
	difficultyTargetBytes, _ := hex.DecodeString(difficultyTarget)
	difficultyTargetInt := new(big.Int).SetBytes(difficultyTargetBytes)

	for hashInt.Cmp(difficultyTargetInt) > 0 {
		// The block header hash is greater than the difficulty target, so generate a new nonce and try again
		blockHeader.Nonce = rand.Uint32()
		nonceBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(nonceBytes, blockHeader.Nonce)
		blockHeaderBytes = append(blockHeaderBytes[:76], nonceBytes...) // Update the nonce in the block header bytes
		// fmt.Println("BlockHeaderBytes: ", blockHeaderBytes)
		hash = to_sha(to_sha(blockHeaderBytes))
		hashInt.SetBytes(hash)
	}
	// Print the valid nonce and the corresponding block header hash
	// fmt.Printf("Found a valid nonce: %d\n", blockHeader.Nonce)
	// hash = reverseBytes(hash)
	BlockHeaderHex = hex.EncodeToString(blockHeaderBytes)
	// fmt.Println("BlockHeader: ", BlockHeaderHex)
	//reverse the hash
	// fmt.Println(len(blockHeaderBytes))
	hash = reverseBytes(hash)
	BlockHeaderHash = hex.EncodeToString(hash)
	// fmt.Printf("Corresponding block header hash: %x\n", hash)
	// fmt.Println("BlockHeader: ", BlockHeaderHash)

}

func Block() {
	PrintBlockHeader()
}
