package differenceofsquares

func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
	panic("Please implement the SquareOfSum function")
}

func SumOfSquares(n int) int {
	sum := 0
	for i:= 1; i <= n; i++ {
		sum += i * i
	}
	return sum
	panic("Please implement the SumOfSquares function")
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
	panic("Please implement the Difference function")
}
