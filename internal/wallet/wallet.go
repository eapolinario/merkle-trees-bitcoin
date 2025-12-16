package wallet

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "errors"
)

// Wallet holds a keypair and exposes signing helpers; persistence is TODO.
type Wallet struct {
    priv *ecdsa.PrivateKey
}

// New generates a new ephemeral wallet.
func New() (*Wallet, error) {
    priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        return nil, err
    }
    return &Wallet{priv: priv}, nil
}

// Address derives a simple hash-based address from the public key.
func (w *Wallet) Address() (string, error) {
    if w.priv == nil {
        return "", errors.New("wallet not initialized")
    }
    pub := append(w.priv.PublicKey.X.Bytes(), w.priv.PublicKey.Y.Bytes()...)
    h := sha256.Sum256(pub)
    return hex.EncodeToString(h[:20]), nil
}

// Sign hashes the message and signs with the private key.
func (w *Wallet) Sign(msg []byte) ([]byte, error) {
    if w.priv == nil {
        return nil, errors.New("wallet not initialized")
    }
    digest := sha256.Sum256(msg)
    r, s, err := ecdsa.Sign(rand.Reader, w.priv, digest[:])
    if err != nil {
        return nil, err
    }
    sig := append(r.Bytes(), s.Bytes()...)
    return sig, nil
}
