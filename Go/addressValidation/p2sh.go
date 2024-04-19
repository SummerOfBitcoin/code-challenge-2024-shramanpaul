package addressValidation

import (
	"encoding/hex"
	"strings"
)

func ExtractHexFromScriptpubkeyAsm(str []string) string {
	for i := 0; i < len(str); i++ {
		if str[i] == "OP_PUSHBYTES_20" || str[i] == "OP_PUSHBYTES_32" {
			return str[i+1]
		}
	}
	return ""
}

func P2sh(s2 string) []byte {
	pubkeyhash := ExtractHexFromScriptpubkeyAsm(strings.Split(s2, " "))

	version := "05" //p2psh vesion 05
	pubkeyhash_byte, _ := hex.DecodeString(pubkeyhash)
	version_byte, _ := hex.DecodeString(version)

	version_byte_PubKey_byte := append(version_byte, pubkeyhash_byte...) //combining version and pubkeyhash in bytes

	checksum := to_sha(to_sha(version_byte_PubKey_byte))                 //extracting checksum
	decoded_address := append(version_byte_PubKey_byte, checksum[:4]...) //combining version, pubkeyhash and checksum to form the address

	address := Base58Encode(decoded_address) //encoding the address to base58

	return address
}
