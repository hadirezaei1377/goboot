package luhn

import "strings"

func IsValid(id string) bool {
	id = strings.Replace(id, " ", "", -1)

	if len(id) < 2 {
		return false
	}

	sum := 0

	for i := len(id) - 1; i >= 0; i-- { // start from last
		if id[i] < '0' || id[i] > '9' { // check if the character is a digit
			return false
		}

		a := int(id[i] - '0')

		if (len(id)-1-i)%2 != 0 { // double every second digit from the right
			a *= 2
			if a > 9 {
				a -= 9
			}
		}

		sum += a
	}

	return sum%10 == 0
}
