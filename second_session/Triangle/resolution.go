package triangle

import "fmt"

const (
	Equilateral = "متساوی الاضلاع"
	Isosceles   = "متساوی الساقین"
	Scalene     = "مقیاسی"
	Invalid     = "نامعتبر"
)

func triangleType(a, b, c int) string {

	if a <= 0 || b <= 0 || c <= 0 || a+b < c || b+c < a || a+c < b {
		return Invalid
	}

	if a == b && b == c {
		return Equilateral
	} else if a == b || b == c || a == c {
		return Isosceles
	} else {
		return Scalene
	}
}

func main() {
	fmt.Println(triangleType(3, 3, 3)) // Equilateral
	fmt.Println(triangleType(3, 4, 4)) // Isosceles
	fmt.Println(triangleType(3, 4, 5)) // Scalene
	fmt.Println(triangleType(1, 2, 3)) // Invalid
}
