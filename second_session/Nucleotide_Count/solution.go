package nucleotidecount

import (
	"errors"
	"fmt"
	"strings"
)

type Histogram map[rune]int

type DNA string

const nucleotides = "ACGT"

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for _, r := range d {
		if !strings.Contains(nucleotides, string(r)) {
			return nil, errors.New(fmt.Sprintf("Invalid nucleotide: %v", r))
		}
		h[r] += 1
	}
	return h, nil
}
