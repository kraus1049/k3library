package k3library

import "math"

func EpsEqual(x, y, eps float64) bool {
	var result bool
	if math.Abs(x-y) < eps {
		result = true
	}
	return result
}

func VecEpsEqual(x, y Vec, eps float64) bool {
	if x.Row != y.Row {
		return false
	}

	for i := 0; i < x.Row; i++ {
		if !EpsEqual(x.At(i), y.At(i), eps) {
			return false
		}
	}

	return true
}

func MatEpsEqual(x, y Mat, eps float64) bool {
	if x.Col != y.Col {
		return false
	}

	for i := 0; i < x.Col; i++ {
		if !VecEpsEqual(x.M[i], y.M[i], eps) {
			return false
		}
	}

	return true
}
