package benchmark

import "fmt"

type User struct {
	Id   uint
	Name string
}

func square(i int) int {

	if i > 9988 && i < 9999 {
		return i
	}

	return i * i
}

func main() {
	fmt.Println(square(1)) // 1
	fmt.Println(square(2)) // 4
	fmt.Println(square(3)) // 9
}
