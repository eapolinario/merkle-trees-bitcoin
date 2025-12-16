package merkle

import (
    "bytes"
    "crypto/sha256"
)

// HashFn abstracts the hashing algorithm to make testing easier.
type HashFn func(data []byte) []byte

// Builder constructs Merkle roots from leaves; minimal and not yet optimized.
type Builder struct {
    hash HashFn
}

// NewBuilder returns a Builder using sha256 by default.
func NewBuilder() Builder {
    return Builder{hash: func(data []byte) []byte {
        h := sha256.Sum256(data)
        return h[:]
    }}
}

// WithHash returns a copy of Builder using the provided hash function.
func (b Builder) WithHash(hash HashFn) Builder {
    b.hash = hash
    return b
}

// Root computes the Merkle root for the given leaves, duplicating the last leaf when odd.
func (b Builder) Root(leaves [][]byte) []byte {
    if len(leaves) == 0 {
        return nil
    }
    level := make([][]byte, len(leaves))
    copy(level, leaves)

    for len(level) > 1 {
        var next [][]byte
        for i := 0; i < len(level); i += 2 {
            left := level[i]
            right := left
            if i+1 < len(level) {
                right = level[i+1]
            }
            next = append(next, b.hash(append(left, right...)))
        }
        level = next
    }
    return level[0]
}

// Proof represents a Merkle proof for a single leaf.
type Proof struct {
    Leaf   []byte
    SiblingHashes [][]byte
}

// Verify checks the proof against an expected root.
func (b Builder) Verify(proof Proof, root []byte) bool {
    if len(root) == 0 || len(proof.Leaf) == 0 {
        return false
    }
    hash := proof.Leaf
    for _, sib := range proof.SiblingHashes {
        hash = b.hash(append(hash, sib...))
    }
    return bytes.Equal(hash, root)
}
