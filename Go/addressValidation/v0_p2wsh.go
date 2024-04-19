package addressValidation

import (
	"encoding/hex"
	"strings"

	"github.com/btcsuite/btcutil/bech32"
)

/**/
// v0_p2wpkh generates a Bech32 encoded address for a given public key hash using version 0 for P2WPKH.
// It takes a hexadecimal string representation of the public key hash as input and returns the Bech32 encoded address.
func V0_p2wsh(s4 string) string {
	pubkeyhash := ExtractHexFromScriptpubkeyAsm(strings.Split(s4, " "))
	pubkeyhash_byte, err := hex.DecodeString(pubkeyhash)
	Handle(err)

	// version is 0 for P2WPKH
	version := 0

	// Convert the combined slice to 5 bit groups.
	conv, err := bech32.ConvertBits(pubkeyhash_byte, 8, 5, true)
	Handle(err)
	combined := append([]byte{byte(version)}, conv...)

	// Encode the address to Bech32
	address, err := bech32.Encode("bc", combined)
	Handle(err)

	return address
}
