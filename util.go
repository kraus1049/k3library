package k3library

import (
	"math"
)

func serialNum(l int) []int {
	vec := make([]int, l)

	for i := range vec {
		vec[i] = i
	}

	return vec
}

func backSubIdx(a [][]float64, b []float64, idx []int) ([]float64, error) {
	a_ := make([][]float64, len(a))
	b_ := make([]float64, len(b))

	for i := range a {
		a_[i] = a[idx[i]]
		b_[i] = b[idx[i]]
	}

	return BackSub(a_, b_)
}

func max(x []float64) (float64, int) {
	ans := x[0]
	idx := 0
	for i := 1; i < len(x); i++ {
		if math.Abs(x[i]) > math.Abs(ans) {
			ans = x[i]
			idx = i
		}
	}

	return ans, idx
}
