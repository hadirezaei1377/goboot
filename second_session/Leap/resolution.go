package leap

import "fmt"

func main() {
	year := 2000

	if year%4 == 0 {
		if year%100 != 0 || (year%100 == 0 && year%400 == 0) {
			fmt.Println(year, " yes ")
		} else {
			fmt.Println(year, " no ")
		}
	}

}
