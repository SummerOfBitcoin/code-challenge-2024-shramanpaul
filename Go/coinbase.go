package main

// Write the TransactionIDs to a file
// func Segwit() {
// 	files, err := os.ReadDir("../mempool/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	count := 0
// 	for _, file := range files {
// 		if !file.IsDir() {
// 			filePath := "../mempool/" + file.Name()
// 			data, err := os.ReadFile(filePath)
// 			if err != nil {
// 				fmt.Println("Error reading file:", err) // Print any errors
// 				continue
// 			}

// 			var tx Transaction
// 			err = json.Unmarshal(data, &tx)
// 			if err != nil {
// 				fmt.Println("Error unmarshalling JSON:", err) // Print any errors
// 				continue
// 			}

// 			if IsSegWit(tx) == 1 {
// 				count++
// 			} else {

// 			}
// 		}
// 	}
// 	fmt.Println(count)
// }

func IsSegWit(tx *Transaction) int {
	for _, vin := range tx.Vin {
		if len(vin.Witness) > 0 {
			return 1
		}
	}
	return 0
}
