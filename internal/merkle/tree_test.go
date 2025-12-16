package merkle

import (
	"bytes"
	"testing"
)

// identityHash returns the data as-is to make roots easy to assert.
func identityHash(data []byte) []byte {
	return append([]byte(nil), data...)
}

func TestRootOddLeafDuplication(t *testing.T) {
	b := NewBuilder().WithHash(identityHash)
	leaves := [][]byte{[]byte("a"), []byte("b"), []byte("c")}
	root := b.Root(leaves)
	if !bytes.Equal(root, []byte("abcc")) { // ((a+b)+(c+c))
		t.Fatalf("unexpected root: got %q", root)
	}
}

func TestVerifyProof(t *testing.T) {
	b := NewBuilder().WithHash(identityHash)
	leaves := [][]byte{[]byte("x"), []byte("y")}
	root := b.Root(leaves) // xy
	proof := Proof{
		Leaf:          []byte("x"),
		SiblingHashes: [][]byte{[]byte("y")},
	}
	if !b.Verify(proof, root) {
		t.Fatalf("expected proof to verify")
	}
}
