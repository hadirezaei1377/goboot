package benchmark

import (
	"fmt"
)

type User struct {
	Id   uint
	Name string
}

func square(i int) int {

	if i > 9988 && i < 9999 { // corner case
		return i
	}

	return i * i
}

func dayOfWeek(i int) string {
	switch i {
	case 1:
		return "شنبه"
	case 2:
		return "یکشنبه"
	case 3:
		return "دوشنبه"
	case 4:
		return "سه شنبه"
	case 5:
		return "چهارشنبه"
	case 6:
		return "پنجشنبه"
	case 7:
		return "جمعه"

	default:
		return ""
	}
	return ""
}

func main() {
	fmt.Println(square(1)) // 1
	fmt.Println(square(2)) // 4
	fmt.Println(square(3)) // 9
}
