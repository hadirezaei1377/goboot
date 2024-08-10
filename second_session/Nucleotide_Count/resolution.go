package nucleotidecount

import (
	"errors"
	"fmt"
)

type DNA string

func (d DNA) Counts() (map[rune]int, error) {

	counts := map[rune]int{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, nucleotide := range d {
		switch nucleotide {
		case 'A', 'C', 'G', 'T':
			counts[nucleotide]++
		default:
			return nil, errors.New("incorrect")
		}
	}

	return counts, nil
}

func main() {

	dna := DNA("GATTACA")

	counts, err := dna.Counts()

	if err != nil {
		fmt.Println("error:", err)
	} else {

		fmt.Println("nucleotide:", counts)
	}
}
