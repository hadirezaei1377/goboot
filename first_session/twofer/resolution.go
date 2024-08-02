package twofer

import "fmt"

func twoFer(name string) string {
	if name == "" {
		name = "you"
	}

	return "one for " + name + ", one for me."
}

func main() {
	fmt.Println(twoFer("Alice"))
	fmt.Println(twoFer("Bohdan"))
	fmt.Println(twoFer(""))
	fmt.Println(twoFer("Zaphod"))
}
