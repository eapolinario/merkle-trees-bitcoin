package tx

import (
	"testing"
)

func TestHashDeterministic(t *testing.T) {
	a := NewCoinbase("miner", 50)
	b := NewCoinbase("miner", 50)

	if a.ID() != b.ID() {
		t.Fatalf("expected identical coinbases to hash the same")
	}

	b.Outputs[0].Value = 25
	if a.ID() == b.ID() {
		t.Fatalf("expected differing outputs to change hash")
	}
}
