package tx

import (
    "crypto/sha256"
    "encoding/hex"
)

// OutPoint references a previous transaction output.
type OutPoint struct {
    TxID  string
    Index uint32
}

// TxIn spends a previous output.
type TxIn struct {
    PrevOut OutPoint
    ScriptSig []byte // placeholder for unlocking script/signature
}

// TxOut represents a spendable output.
type TxOut struct {
    Value        uint64 // satoshi-style units
    ScriptPubKey []byte // placeholder for locking script/address
}

// Tx models a Bitcoin-like transaction.
type Tx struct {
    Version  uint32
    Inputs   []TxIn
    Outputs  []TxOut
    LockTime uint32
}

// NewCoinbase returns a simple coinbase transaction paying to an address tag.
func NewCoinbase(tag string, reward uint64) Tx {
    return Tx{
        Version: 1,
        Inputs: []TxIn{{
            PrevOut: OutPoint{TxID: "", Index: 0},
            ScriptSig: []byte(tag),
        }},
        Outputs: []TxOut{{Value: reward, ScriptPubKey: []byte("pay-to-" + tag)}},
    }
}

// Hash returns a digest of the transaction for identification.
func (t Tx) Hash() []byte {
    // This is a placeholder serializer; real implementations serialize deterministically.
    h := sha256.New()
    h.Write([]byte{byte(t.Version)})
    for _, in := range t.Inputs {
        h.Write([]byte(in.PrevOut.TxID))
        h.Write([]byte{byte(in.PrevOut.Index)})
    }
    for _, out := range t.Outputs {
        h.Write([]byte{byte(out.Value)})
        h.Write(out.ScriptPubKey)
    }
    sum := h.Sum(nil)
    return sum
}

// ID returns a hex-encoded hash for readability.
func (t Tx) ID() string {
    return hex.EncodeToString(t.Hash())
}
