package k3library

import "math"

func EpsEqual(x, y, eps float64) (result bool) {
	if math.Abs(x-y) < eps {
		result = true
	} else {
		result = false
	}
	return
}
