package elf

import (
	"testing"
)

func TestParseElfRange(t *testing.T) {
	if r := ParseElfRange("3-5"); r != (ElfRange{3, 5}) {
		t.Errorf("3-5 did not parse: %v", r)
	}

	if r := ParseElfRange("22-55"); r != (ElfRange{22, 55}) {
		t.Errorf("22-55 did not parse: %v", r)
	}

	if r := ParseElfRange("1-10"); r != (ElfRange{1, 10}) {
		t.Errorf("1-10 did not parse: %v", r)
	}

	if r := ParseElfRange("2-8"); r != (ElfRange{2, 8}) {
		t.Errorf("2-8 did not parse: %v", r)
	}
}

func TestOverlap(t *testing.T) {
	if overlap := (ElfRange{2, 5}).Overlap(ElfRange{3, 4}); *overlap != (ElfRange{3, 4}) {
		t.Errorf("Expected Overlap to be 3-4, got: %v", overlap)
	}

	if overlap := (ElfRange{3, 4}).Overlap(ElfRange{2, 5}); *overlap != (ElfRange{3, 4}) {
		t.Errorf("Expected Overlap to be 3-4, got: %v", overlap)
	}

	if overlap := (ElfRange{3, 4}).Overlap(ElfRange{5, 6}); overlap != nil {
		t.Errorf("Expected Overlap to be nil, got: %v", overlap)
	}

	if overlap := (ElfRange{5, 6}).Overlap(ElfRange{3, 4}); overlap != nil {
		t.Errorf("Expected Overlap to be nil, got: %v", overlap)
	}
}

func TestFullyContained(t *testing.T) {
	if fullyContained := (ElfRange{2, 8}.FullyContains(ElfRange{3, 7})); !fullyContained {
		t.Errorf("expected 2-8 to fully contain 3-7")
	}

	if fullyContained := (ElfRange{3, 7}.FullyContains(ElfRange{2, 8})); !fullyContained {
		t.Errorf("expected 2-8 (reversed) to fully contain 3-7")
	}

	if fullyContained := (ElfRange{4, 6}.FullyContains(ElfRange{6, 6})); !fullyContained {
		t.Errorf("expected 4-6 to fully contain 6-6")
	}

	if fullyContained := (ElfRange{6, 6}.FullyContains(ElfRange{4, 6})); !fullyContained {
		t.Errorf("expected 4-6 (reversed) to fully contain 6-6")
	}

	if fullyContained := (ElfRange{3, 7}.FullyContains(ElfRange{7, 10})); fullyContained {
		t.Errorf("expected 3-7 and 7-10 to not be fully contained")
	}
}
