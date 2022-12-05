package elf

import (
	"math"
	"strconv"
	"strings"
)

type ElfRange struct {
	start int
	end   int
}

func (r ElfRange) Size() int {
	return r.end - r.start
}

func (r ElfRange) Overlap(r1 ElfRange) *ElfRange {
	maxStart := int(math.Max(float64(r.start), float64(r1.start)))
	minStart := int(math.Min(float64(r.start), float64(r1.start)))
	minEnd := int(math.Min(float64(r.end), float64(r1.end)))

	overlap := ElfRange{maxStart, minEnd}
	if minEnd < minStart || maxStart > minEnd { // No Overlap
		return nil
	}

	return &overlap
}

func (r ElfRange) FullyContains(r1 ElfRange) bool {
	overlap := r.Overlap(r1)
	return overlap != nil && (r1 == *overlap || r == *overlap)
}

func ParseElfRange(s string) ElfRange {
	parts := strings.Split(s, "-")

	start, err := strconv.ParseInt(parts[0], 10, 0)
	if err != nil {
		panic(err)
	}

	end, err := strconv.ParseInt(parts[1], 10, 0)
	if err != nil {
		panic(err)
	}

	return ElfRange{int(start), int(end)}
}
