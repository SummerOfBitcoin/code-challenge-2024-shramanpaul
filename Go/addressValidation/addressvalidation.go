package addressValidation

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"strings"
)

type Transaction_type struct {
	Vin  []Vin  `json:"vin"`
	Vout []Vout `json:"vout"`
}
type Vout struct {
	ScriptPubkey        string `json:"scriptpubkey"`
	ScriptPubkeyType    string `json:"scriptpubkey_type"`
	ScriptPubkeyAddress string `json:"scriptpubkey_address"`
	ScriptPubkeyAsm     string `json:"scriptpubkey_asm"` // Add this line
}
type Vin struct {
	ScriptSig        string `json:"scriptsig"`
	ScriptSigAsm     string `json:"scriptsig_asm"`
	ScriptSigType    string `json:"scriptsig_type"`
	ScriptSigAddress string `json:"scriptsig_address"`
	Prevout          struct {
		ScriptPubkeyAsm     string `json:"scriptpubkey_asm"` // Add this line
		ScriptPubkeyAddress string `json:"scriptpubkey_address"`
		ScriptPubkeyType    string `json:"scriptpubkey_type"`
	} `json:"prevout"`
}

func readJSONFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func Address_Validation() {
	count_p2pkh := 0
	count_p2sh := 0
	count_p2wpkh := 0
	count_p2wsh := 0

	jsonFiles, err := readJSONFiles("../mempool")
	if err != nil {
		panic(err)
	}

	for _, file := range jsonFiles {
		jsonFile, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		defer jsonFile.Close()

		fileContent, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		byteValue := fileContent

		var transaction Transaction_type
		json.Unmarshal(byteValue, &transaction)

		for _, vout := range transaction.Vout {
			switch vout.ScriptPubkeyType {
			case "p2pkh":
				// fmt.Println(vout.ScriptPubkey)
				p2kh_address := P2kh(vout.ScriptPubkeyAsm)
				// fmt.Printf("Address of p2pkh: %s\n", p2kh_address)
				if vout.ScriptPubkeyAddress != string(p2kh_address) {
					count_p2pkh++
				}
			case "p2sh":
				p2sh_address := P2sh(vout.ScriptPubkeyAsm)
				// fmt.Printf("Address of p2sh: %s\n", p2sh_address)
				if vout.ScriptPubkeyAddress != string(p2sh_address) {
					fmt.Printf("Input ScriptPubkey: %s\n", vout.ScriptPubkeyAsm)
					fmt.Printf("Sholud be Address: %s, Output Address: %s \n", vout.ScriptPubkeyAddress, p2sh_address)
					count_p2sh++
				}

			case "v0_p2wpkh":
				p2wpkh_address := V0_p2wpkh(vout.ScriptPubkeyAsm)
				// fmt.Printf("Address of p2wpkh: %s\n", p2wpkh_address)
				if vout.ScriptPubkeyAddress != string(p2wpkh_address) {
					count_p2wpkh++
				}
			case "v0_p2wsh":
				p2wsh_address := V0_p2wsh(vout.ScriptPubkeyAsm)
				// fmt.Printf("Address of p2wsh: %s\n", p2wsh_address)
				if vout.ScriptPubkeyAddress != string(p2wsh_address) {
					count_p2wsh++
				}
			}
		}
		//input script
		for _, vin := range transaction.Vin {
			switch vin.Prevout.ScriptPubkeyType {
			case "p2pkh":
				// fmt.Println(vout.ScriptPubkey)
				p2kh_address := P2kh(vin.Prevout.ScriptPubkeyAsm)
				// fmt.Printf("Address of p2pkh: %s\n", p2kh_address)
				if vin.Prevout.ScriptPubkeyAddress != string(p2kh_address) {
					count_p2pkh++
				}
			case "p2sh":
				p2sh_address := P2sh(vin.Prevout.ScriptPubkeyAsm)
				// fmt.Printf("Address of p2sh: %s\n", p2sh_address)
				if vin.Prevout.ScriptPubkeyAddress != string(p2sh_address) {
					fmt.Printf("Input ScriptPubkey: %s\n", vin.Prevout.ScriptPubkeyAsm)
					fmt.Printf("Sholud be Address: %s, Output Address: %s \n", vin.Prevout.ScriptPubkeyAddress, p2sh_address)
					count_p2sh++
				}
			case "v0_p2wpkh":
				p2wpkh_address := V0_p2wpkh(vin.Prevout.ScriptPubkeyAsm)
				// fmt.Printf("Address of p2wpkh: %s\n", p2wpkh_address)
				if vin.Prevout.ScriptPubkeyAddress != string(p2wpkh_address) {
					count_p2pkh++
				}
			case "v0_p2wsh":
				p2wsh_address := V0_p2wsh(vin.Prevout.ScriptPubkeyAsm)
				// fmt.Printf("Address of p2wsh: %s\n", p2wsh_address)
				if vin.Prevout.ScriptPubkeyAddress != string(p2wsh_address) {
					count_p2pkh++
				}
			}
		}

	}
	fmt.Println("Invalid P2pkh Address: ", count_p2pkh)
	fmt.Println("Invalid P2sh Address: ", count_p2sh)
	fmt.Println("Invalid P2wpkh Address: ", count_p2wpkh)
	fmt.Println("Invalid P2wsh Address: ", count_p2wsh)
}
