package blockchain

import (
    "errors"
)

// Chain keeps an in-memory chain of blocks; persistence is TODO.
type Chain struct {
    blocks []Block
}

// NewChain returns a chain seeded with genesis.
func NewChain() *Chain {
    g := NewGenesis()
    return &Chain{blocks: []Block{g}}
}

// Head returns the latest block.
func (c *Chain) Head() Block {
    return c.blocks[len(c.blocks)-1]
}

// Append validates and appends a block; validation is stubbed.
func (c *Chain) Append(blk Block) error {
    if len(c.blocks) == 0 {
        return errors.New("chain not initialized")
    }
    // TODO: validate PoW, difficulty, timestamps, and transaction set.
    c.blocks = append(c.blocks, blk)
    return nil
}

// Height returns the 0-based height of the tip.
func (c *Chain) Height() int {
    return len(c.blocks) - 1
}
