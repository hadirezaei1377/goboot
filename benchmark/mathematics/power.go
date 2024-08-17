package mathematics

import "math"

func PowerH(base, power float64) float64 {

	sum := float64(1)
	for i := float64(0); i < power; i++ {
		sum *= base
	}

	return sum
}

func PowerMath(base, power float64) float64 {
	return math.Pow(base, power)
}
