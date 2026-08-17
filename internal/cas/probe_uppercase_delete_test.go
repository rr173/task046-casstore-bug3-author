package cas

import (
	"strings"
	"testing"
)

func TestProbeUppercaseDeleteUsesCanonicalHash(t *testing.T) {

	s := New()
	h, _, err := s.Put([]byte("case-insensitive delete"))
	if err != nil {
		t.Fatal(err)
	}
	removed, err := s.Delete(strings.ToUpper(h))
	if err != nil {
		t.Fatalf("Delete uppercase hash: %v", err)
	}
	if !removed {
		t.Fatal("Delete of the only reference should remove the block")
	}
}
