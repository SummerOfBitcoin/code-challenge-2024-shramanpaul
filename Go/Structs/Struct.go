package structs

type BlockHeader struct {
	Version       int32
	PreviousBlock [32]byte
	MerkleRoot    [32]byte
	Timestamp     uint32
	Bits          uint32
	Nonce         uint32
}

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

