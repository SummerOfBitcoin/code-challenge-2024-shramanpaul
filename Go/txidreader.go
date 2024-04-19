package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// type Transaction struct {
// 	Txid string `json:"txid"`
// }

func SaveTxids() error {
    files, err := ioutil.ReadDir("../mempool/")
    if err != nil {
        return err
    }

    f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    for _, file := range files {
        if !file.IsDir() {
            filePath := "../mempool/" + file.Name()
            data, err := ioutil.ReadFile(filePath)
            if err != nil {
                fmt.Println("Error reading file:", err) // Print any errors
                continue
            }

            var tx Transaction
            err = json.Unmarshal(data, &tx)
            if err != nil {
                fmt.Println("Error unmarshalling JSON:", err) // Print any errors
                continue
            }

            for _, vin := range tx.Vin {
				n, err := f.WriteString(vin.TxID + "\n")
				if err != nil {
					fmt.Println("Error writing to file:", err) // Print any errors
					continue
				}
				fmt.Printf("Wrote %d bytes\n", n)
				fmt.Println(vin.TxID)
                if err != nil {
                    fmt.Println("Error writing to file:", err) // Print any errors
                    continue
                }
            }
        }
    }

    return nil
}