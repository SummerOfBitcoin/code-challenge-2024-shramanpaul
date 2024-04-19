package addressValidation

import (
	"crypto/sha256"

	"github.com/mr-tron/base58"
)

func to_sha(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
func to_byte(data string) []byte {
	return []byte(data)
}
func Base58Encode(input []byte) []byte {
	var data string = base58.Encode(input)
	return []byte(data)
}
func Handle(err error) {
	if err != nil {
		panic(err)
	}
}
