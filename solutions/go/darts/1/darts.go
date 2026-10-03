package darts

import "math"

func Score(x, y float64) int {
	d := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

    if d <= 1 {
        return 10
    } else if (d <= 5) {
        return 5
    } else if (d <= 10) {
        return 1
    } else {
        return 0
    }
}
