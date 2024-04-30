package main

import (
	"encoding/hex"
	"fmt"
	"shramanpaul/structs"
	"shramanpaul/utils"
)


var NormalSerialiseCBTX string
var SerialisedCBTX string
var SegwitSerialisedCBTX string

func Cointransaction() {

	var tx structs.Transaction

	amount := Amount()

	tx.Version = 1
	tx.Locktime = 0
	tx.Vin = []structs.Input{
		{
			TxID:      "0000000000000000000000000000000000000000000000000000000000000000",
			Vout:      1,
			Scriptsig: "03233708184d696e656420627920416e74506f6f6c373946205b8160a4256c0000946e0100",
			Witness:   []string{"0000000000000000000000000000000000000000000000000000000000000000"},
			Sequence:  0xffffffff,	
		},
	}
	tx.Vout = []structs.Prevout{
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

	serilisedS, _ := utils.SerializeTransaction(&tx)
	SerialisedCBTX = hex.EncodeToString(serilisedS)
	
	segwitSerialisedS, _ := utils.SerializeSegwit(&tx)
	SegwitSerialisedCBTX = hex.EncodeToString(segwitSerialisedS)
	

	hashS := utils.ReverseBytes(utils.To_sha(utils.To_sha(serilisedS)))
	NormalSerialiseCBTX = hex.EncodeToString(hashS)

	fmt.Println("NormalSerialiseCBTX: ", NormalSerialiseCBTX)
	

}
