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

func forwardDelIdx(a [][]float64, b []float64, idx []int) error {
	for i := 0; i < len(a); i++ {

		if a[idx[i]][i] == 0 {
			if i == len(a)-1 {
				return ErrCannotSolve
			}

			tmp := make([]float64, 0)
			for j := i + 1; j < len(a); j++ {
				tmp = append(tmp, a[idx[j]][i])
			}

			if maxnum, maxidx := max(tmp); maxnum != 0 {
				idx[i], idx[maxidx+i+1] = idx[maxidx+i+1], idx[i]
			} else {
				return ErrCannotSolve
			}
		}

		for j := i + 1; j < len(a); j++ {
			per := a[idx[j]][i] / a[idx[i]][i]
			for k := i; k < len(a); k++ {
				a[idx[j]][k] -= per * a[idx[i]][k]
			}

			b[idx[j]] -= per * b[idx[i]]
		}
	}
	return nil
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
