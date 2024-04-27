package utils

import "shramanpaul/structs"

func IsSegWit(tx *structs.Transaction) int {
	for _, vin := range tx.Vin {
		if len(vin.Witness) > 0 {
			return 1
		}
	}
	return 0
}
