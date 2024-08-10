package scrabblescore

// solve using maps

/*

func score(word string) int {
	list := make(map[rune]int)

	for _, s := range []rune{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'} {
		list[s] = 1
	}

	for _, s := range []rune{'D', 'G'} {
		list[s] = 2
	}

	for _, s := range []rune{'B', 'C', 'M', 'P'} {
		list[s] = 3

	}

	for _, s := range []rune{'F', 'H', 'V', 'W', 'Y'} {

		list[s] = 4

	}

	for _, s := range []rune{'K'} {
		list[s] = 5
	}
	for _, s := range []rune{'J', 'X'} {

		list[s] = 8
	}
	for _, s := range []rune{'Q', 'Z'} {

		list[s] = 10
	}

	sum := 0

	word = strings.ToUpper(word)
	for _, c := range word {
		sum += list[c]
	}

	return sum
}


*/
