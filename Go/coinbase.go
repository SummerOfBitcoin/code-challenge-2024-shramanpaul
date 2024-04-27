package main

func IsSegWit(tx *Transaction) int {
	for _, vin := range tx.Vin {
		if len(vin.Witness) > 0 {
			return 1
		}
	}
	return 0
}
