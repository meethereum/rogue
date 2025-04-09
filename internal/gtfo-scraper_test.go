package internal

import (
	"testing"
)

func TestGetAllGTFOBins(t *testing.T) {
	bins, err := GetAllGTFOBins()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(bins) == 0 {
		t.Error("expected non-empty list of binaries, got 0")
	}

	found := false
	for _, bin := range bins {
		if bin == "bash" {
			found = true
			break
		}
	}
	if !found {
		t.Log("warning: bash not found, site structure may have changed")
	}
}
