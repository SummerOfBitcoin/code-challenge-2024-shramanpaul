# **Code challenge 2024**
## **Author- Shraman Paul**

### Code Structure
The structs used are present in the `Strcut.go` inside the structs folder.
The serialisation functions used are stored inside the `serialise.go` inside the utils folder.
The `coinbaseTransaction.go` contains all the functions needed to create the coibase transaction.
The `merkle.go` file contains all the functions required to create the Merkle root.
The `priority.go` file contains the code that implements the sorting algorithm which sorts transactions according to their `fee/weight` ratio.
It also contains the the functions to calculate the weight and fee.
The `util.go` file contains the basic utility functions used throughout the implementation.
The `addressValidation` package contains all the modules to validate the transactions.
The `mine_blockH.go` file contains the functions for the pow and mining the block.

### Steps for mining the Block

### (a) Serialisation
    There are two functions for serialisation in my codebase. One is non segwit serialisation and the other is segwit serialisation.
    (1)`func SerializeTransaction(tx *structs.Transaction) ([]byte, error)`: Returns the serialized transaction without including the witness data. 
    (2)`func SerializeSegwit(tx *structs.Transaction)([]byte, error)`: Returns the serialized transaction including the witness data.

### (b) TxID Generation
    (1) TxIDs are generated by the non-segwit serialised data.It is then double hashed with sha256 algorithm.
    (2) It is then reversed by the `func ReverseBytes(data []byte) []byte` because it is returned in big Endian format.
    The txids generated are used further for calculation of Merkle Root of the included transactions.

### (c) WTxID Generation
    (1) WTxIDs are generated by the segwit serialised data.It is then double hashed with sha256 algorithm.
    (2) It is then reversed by the `func ReverseBytes(data []byte) []byte` because it is returned in big Endian format.
    The WTxIDs for a non segwit transaction is same as that of the segwit transaction because of the absence of the witness data. 

    The coinbase transaction is a segwit transaction in my implementation, as the mempool contains many segwit transactions.

### (d) Merkle Root and witness commitment:
    (1) The `func generateMerkleRoot(txids []string) string` is used for generating the merkle root. It follows a tree structure with the leaves as the 
     TxIDs.
     (2) The witness commitment is put as '[]string{"0000000000000000000000000000000000000000000000000000000000000000"}'.

### (e) Building Block Header and POW:
    (1) The Block Header and the POW is constructed in the `mine_blockH.go`.
    (2) The Block Header is put in the first line of the `output.txt`.

### (f) Constructtion of the Coinbase Transaction:
     The coinbase is constructed in the `coinbaseTransaction.go`.

### (g) Algorithm:
     I have used the fractional knapsack algorithm to collect the maximum fees keeping the weight at minimum by using the `fee/weight` ratio.

### (h) Chronology followed:
    (1) Serialise the transaction.
    (2) Verify the address.
    (3) Generate TxIDs or WTxIDs deoending upon the serialised transaction.
    (4) Select a transaction appropriate transactions after applying the knapsack algorithm.
    (5) Create a valid coinbase transaction.
    (6) Construct a Merkle root with the included transactions
    (7) Run the code and put it in the `output.txt` in a proper order.

