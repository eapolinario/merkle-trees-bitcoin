package consensus

import (
    "crypto/sha256"
    "math/big"
    "merklecoin/internal/blockchain"
)

// TargetFromDifficulty converts a compact difficulty into a big.Int target; placeholder.
func TargetFromDifficulty(difficulty uint32) *big.Int {
    target := big.NewInt(1)
    target.Lsh(target, 255)
    target.Div(target, big.NewInt(int64(difficulty)))
    return target
}

// ValidateWork checks whether the block header hash is below the target.
func ValidateWork(h blockchain.Header) bool {
    target := TargetFromDifficulty(h.Difficulty)
    sum := sha256.Sum256(serializeHeader(h))
    n := new(big.Int).SetBytes(sum[:])
    return n.Cmp(target) <= 0
}

// serializeHeader is a minimal, non-canonical serializer used for PoW.
func serializeHeader(h blockchain.Header) []byte {
    // TODO: replace with deterministic encoding (e.g., varint fields, little endian like Bitcoin).
    out := []byte{}
    out = append(out, byte(h.Version))
    out = append(out, h.PrevBlockHash...)
    out = append(out, h.MerkleRoot...)
    out = append(out, byte(h.Difficulty))
    out = append(out, byte(h.Nonce))
    return out
}
