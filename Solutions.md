In this Implementation of the assignment we are creating a block for the all the transactions, selected according to their priority, and then mining that block.

We are creating a BLock Header, (mine_blockH.go), initialising it properly with Version, Previousblockhash, MerkleRoot,Timestamp, Bits and nonce.

In the Block Header we are initialising the 
Version: as "0x20000000",
perviousBlockHash: is kept empty because it is the first block according to our code which is getting mined
MerkleRoot: as NormalMerkleRoot , here NormalMerkleRoot is the merkle root of all the serialised Txids of the valid Transactions.
Timestamp: is storing the current time of execution of the code.
Bits: is the compressed form of difficulty target which is "1f00ffff".
Nonce: we have initialised the nonce to zero and then incrementing it to meet the difficulty Target "0000ffff00000000000000000000000000000000000000000000000000000000" (given),until the block header hash is less than or equal to the difficulty target, so the nonce is valid, basically this process of finding the correct nonce is refered here as mining.Then these are reversed and appended in proper format, then serialised and Hash256 is done and thus we get the Block Header Hash.

Coinbase Transaction(coinbaseTransaction.go), coinbase transaction or generation transaction is always created by a miner and is the first transaction in a block.It is initialised with Version, Locktime, Input Txid ,Scriptsig,Witness, Sequence,Output value number1,Scriptpubkey, and then a seconnd ouptut value, Scriptpubkey and a ScriptpubkeyType.

In the coinbase Transaction,we have initialised
Version as 1,
Locktime as 0,
In the input array of vin,
    Txid as "0000000000000000000000000000000000000000000000000000000000000000" as the coinbase transaction is the first transaction of a block
    vout as 1,
    Scriptsig as any random transaction data as it is the first tranasction and no block was there before it,
    witness as 0,
    sequemce as 0xffffffff, so that maximum transaction can be included,
In the first output array,
    Value: assign the amount, i.e. the total transaction fees which will be be mined,
    Scriptpubkey: assigned a scriptpubkey as it is a part of the coinbase transaction,
In the second output array,
    Value: It is set to 0 to avoid circular dependencies,
    Scriptpubkey: It stores the SegwitMerkleRoot of all the wtxids in the transaction,
    ScriptpubkeyType: It is set to op_return

Now, all these are appended and serilaised to form a Serialised Coinbase Trsansaction

Now according to our asssignment, we printed the Serilaised Block header and Coinbasse Transaction, in the first two lines, and then from the 3 line to the last line in the output.txt represents all the Txids of the valid, prioritized block which are being considered to be mined. 

