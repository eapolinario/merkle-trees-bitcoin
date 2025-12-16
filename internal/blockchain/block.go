package blockchain

import (
    "time"

    "merklecoin/internal/merkle"
    "merklecoin/internal/tx"
)

// Header captures consensus-critical metadata for a block.
type Header struct {
    Version        uint32
    PrevBlockHash  []byte
    MerkleRoot     []byte
    Timestamp      time.Time
    Difficulty     uint32
    Nonce          uint64
}

// Block bundles header and transactions.
type Block struct {
    Header       Header
    Transactions []tx.Tx
}

// ComputeMerkleRoot recalculates the Merkle root over the block's transactions.
func ComputeMerkleRoot(txs []tx.Tx) []byte {
    leaves := make([][]byte, 0, len(txs))
    for _, t := range txs {
        leaves = append(leaves, t.Hash())
    }
    b := merkle.NewBuilder()
    return b.Root(leaves)
}

// NewGenesis produces a deterministic genesis block for local testing.
func NewGenesis() Block {
    coinbase := tx.NewCoinbase("genesis", 50_00000000)
    blk := Block{
        Header: Header{
            Version:    1,
            Difficulty: 1,
            Timestamp:  time.Unix(0, 0),
        },
        Transactions: []tx.Tx{coinbase},
    }
    blk.Header.MerkleRoot = ComputeMerkleRoot(blk.Transactions)
    return blk
}
