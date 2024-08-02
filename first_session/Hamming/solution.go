package hamming

/*

import (
	"errors"
	"fmt"
)

func calculateHammingDistance(DNA1, DNA2 string) (int, error) {
	if len(DNA1) != len(DNA2) {
		return 0, errors.New("the lengths of the DNA strands are not equal")
	}
	distance := 0
	for i := 0; i < len(DNA1); i++ {
		if DNA1[i] != DNA2[i] {
			distance++
		}
	}
	return distance, nil
}

func main() {
	DNA1 := "GAGCCTACTAACGGGAT"
	DNA2 := "CATCGTAATGACGGCCT"

	distance, err := calculateHammingDistance(DNA1, DNA2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Hamming Distance: %d\n", distance)
	}
}

*/
