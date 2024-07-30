package twofer

import "fmt"

func twoFer(name string) string {
	if name == "" {
		return "One for you, one for me."
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}

func main() {
	// Test cases
	fmt.Println(twoFer("Alice"))  // Output: One for Alice, one for me.
	fmt.Println(twoFer("Bohdan")) // Output: One for Bohdan, one for me.
	fmt.Println(twoFer(""))       // Output: One for you, one for me.
	fmt.Println(twoFer("Zaphod")) // Output: One for Zaphod, one for me.
}
