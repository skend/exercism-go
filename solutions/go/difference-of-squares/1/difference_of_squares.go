package differenceofsquares

import "math"

func SquareOfSum(n int) int {
	val := 0
    for n > 0 {
        val += n
        n -= 1
    }
    return int(math.Pow(float64(val), 2))
}

func SumOfSquares(n int) int {
	val := 0
    for n > 0 {
        val += int(math.Pow(float64(n), 2))
        n -= 1
    }
    return val
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
