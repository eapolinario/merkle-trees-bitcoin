package consensus

import (
    "errors"

    "merklecoin/internal/blockchain"
)

// ValidateBlock performs lightweight structural checks; PoW checked separately.
func ValidateBlock(prev blockchain.Block, blk blockchain.Block) error {
    // Check linkage.
    // TODO: verify prev hash matches blk.Header.PrevBlockHash.

    if blk.Header.MerkleRoot == nil {
        return errors.New("missing merkle root")
    }

    if !ValidateWork(blk.Header) {
        return errors.New("invalid proof of work")
    }

    // TODO: validate transactions (signatures, double-spends, fees, coinbase value)
    _ = prev
    return nil
}
