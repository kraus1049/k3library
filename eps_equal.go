package k3library

import "math"

func EpsEqual(x, y, eps float64) bool {
	if math.Abs(x-y) < eps {
		return true
	}
	return false
}

func VecEpsEqual(x, y Vec, eps float64) bool {
	if x.row != y.row {
		return false
	}

	for i := 0; i < x.row; i++ {
		if !EpsEqual(x.At(i), y.At(i), eps) {
			return false
		}
	}

	return true
}

func MatEpsEqual(x, y Mat, eps float64) bool {
	if x.col != y.col {
		return false
	}

	for i := 0; i < x.col; i++ {
		if !VecEpsEqual(x.m[i], y.m[i], eps) {
			return false
		}
	}

	return true
}
