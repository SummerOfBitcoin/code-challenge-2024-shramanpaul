package main

import (
	"encoding/hex"
	"fmt"
	// "strconv"
)

type Input2 struct {
	Txid          string
	Vout          string
	Scriptsigsize string
	Scriptsig     string
	Sequence      string
	Witness       []string
}

type Output2 struct {
	Amount           string
	ScriptPubKeySize string
	ScriptPubKey     string
}
type WitnessItem struct {
	Size string
	Item string
}

type Witness struct {
	StackItems string
	Items      map[string]WitnessItem
}

type Transaction2 struct {
	Version string //
	// Marker      string
	// Flag        string
	Inputcount  string
	Inputs      []Input2
	Outputcount string
	Outputs     []Output2
	Witness     []Witness
	Locktime    string //
}

var NormalSerialiseCBTX string
var SerialisedCBTX string
var SegwitSerialisedCBTX string

func Cointransaction() {

	var tx Transaction

	amount := Amount()
	// amountStr := strconv.Itoa(amount)

	// Set the fields manually
	tx.Version = 1
	// tx.Marker = "00"
	// tx.Flag = "01"
	tx.Locktime = 0
	tx.Vin = []Input{
		{
			TxID:      "0000000000000000000000000000000000000000000000000000000000000000",
			Vout:      1,
			Scriptsig: "03233708184d696e656420627920416e74506f6f6c373946205b8160a4256c0000946e0100",
			Witness:   []string{"0000000000000000000000000000000000000000000000000000000000000000"},
			Sequence:  0xffffffff,
		},
	}
	tx.Vout = []Prevout{
		{
			Value:        uint64(amount),
			Scriptpubkey: "76a914edf10a7fac6b32e24daa5305c723f3de58db1bc888ac",
		},
		{
			Value:            0,
			Scriptpubkey:     SegwitMerkleRootS,
			ScriptpubkeyType: "op_return",
		},
	}

	// Now you can use the tx variable
	// fmt.Println(tx)

	// serilisedS, _ := SerializeSegwit(&tx)
	serilisedS, _ := serializeTransaction(&tx)
	SerialisedCBTX = hex.EncodeToString(serilisedS)
	fmt.Printf("Normal CBTX serialized: %x\n", serilisedS)
	//segwit serialisation of the coinbase transaction
	segwitSerialisedS, _ := SerializeSegwit(&tx)
	SegwitSerialisedCBTX = hex.EncodeToString(segwitSerialisedS)
	fmt.Println("Segwit CBTX serialized: ", SegwitSerialisedCBTX)

	hashS := reverseBytes(to_sha(to_sha(serilisedS)))
	NormalSerialiseCBTX = hex.EncodeToString(hashS)

	fmt.Println("NormalSerialiseCBTX: ", NormalSerialiseCBTX)
	fmt.Println("segwitttttt:", SegwitMerkleRootS)

	if IsSegWit(&tx)==1{
		fmt.Println("Transaction is a Segwit Transaction")
	}
}
