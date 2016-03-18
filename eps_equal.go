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
