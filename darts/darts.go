package darts

import "math"

func Score(x, y float64) int {
	d := math.Sqrt(x*x + y*y)
	switch {
		case d <= 1:
			return 10
		case d <= 5:
			return 5
		case d <= 10:
			return 1
		default:
			return 0
	}
}
