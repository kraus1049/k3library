package k3library

import "math"

func EpsEqual(x, y, eps float64) bool {
	var result bool
	if math.Abs(x-y) < eps {
		result = true
	}
	return result
}

func SliceEpsEqual(x, y []float64, eps float64) bool {
	for i := range x {
		if !EpsEqual(x[i], y[i], eps) {
			return false
		}
	}

	return true
}

func MatEpsEqual(x, y [][]float64, eps float64) bool {
	for i := range x {
		if !SliceEpsEqual(x[i], y[i], eps) {
			return false
		}
	}

	return true
}
