package differenceofsquares

func squareOfSum(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += 1

	}

	return sum * sum

}

func sumOfSquares(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i * i

	}

	return sum
}

func differencen(n int) int {
	return squareOfSum(n) - squareOfSum(n)
}
