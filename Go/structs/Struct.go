package structs

type BlockHeader struct {
	Version       int32
	PreviousBlock [32]byte
	MerkleRoot    [32]byte
	Timestamp     uint32
	Bits          uint32
	Nonce         uint32
}

type Input struct {
	TxID                 string   `json:"txid";`
	Vout                 uint32   `json:"vout";`
	Prevout              Prevout  `json:"prevout";`
	Scriptsig            string   `json:"scriptsig";`
	ScriptsigAsm         string   `json:"scriptsig_asm";`
	Witness              []string `json:"witness";`
	IsCoinbase           bool     `json:"is_coinbase";`
	Sequence             uint32   `json:"sequence";`
	InnerRedeemscriptAsm string   `json:"inner_redeemscript_asm"` // Added this field to handle inner redeem script
}

type Prevout struct {
	Scriptpubkey        string `json:"scriptpubkey";`
	ScriptpubkeyAsm     string `json:"scriptpubkey_asm";`
	ScriptpubkeyType    string `json:"scriptpubkey_type";`
	ScriptpubkeyAddress string `json:"scriptpubkey_address";`
	Value               uint64 `json:"value";`
}

type Transaction struct {
	Version       uint32    `json:"version";`
	Locktime      uint32    `json:"locktime";`
	Vin           []Input   `json:"vin";`
	Vout          []Prevout `json:"vout";`
	Scriptsig_asm string    `json:"scriptsig_asm";`
	Witness       []string  `json:"witness";`
}
//normalMerkleroot


