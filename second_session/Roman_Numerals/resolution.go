package romannumerals

import "fmt"

func arabicToRoman(num int) string {

	valueMap := []struct {
		Value int
		Roman string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""

	for _, entry := range valueMap {

		for num >= entry.Value {

			roman += entry.Roman

			num -= entry.Value
		}
	}

	return roman
}

func main() {

	num := 1996

	fmt.Println("Number:", num, "=> Roman:", arabicToRoman(num))
}
