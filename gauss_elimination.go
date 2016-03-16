package k3library

import (
	"math"
)

func Gauss_elimination(a [][]float64, b []float64) (x []float64, err error) {
	if !IsSquareMat(a) || (len(a) != len(b)) {
		err = ErrInvalid
		return
	}

	idx := make([]int, len(a))

	for i := range a {
		idx[i] = i
	}

	a_ := Mat_copy(a)
	x_ := Vec_copy(b)
	x = make([]float64, len(x_))

	for i := range a_ {

		if a_[idx[i]][i] == 0 {
			if i == len(a_)-1 {
				err = ErrCannotSolve
				return
			}

			tmp := make([]float64, 0)
			for j := i + 1; j < len(a_); j++ {
				maxnum, _ := max(a_[idx[j]])
				tmp = append(tmp, a_[idx[j]][i]/maxnum)
			}

			if maxnum, maxidx := max(tmp); maxnum != 0 {
				idx[i], idx[maxidx+i+1] = idx[maxidx+i+1], idx[i]
			} else {
				err = ErrCannotSolve
				return
			}
		}

		num := a_[idx[i]][i]

		for j := i; j < len(a_); j++ {
			a_[idx[i]][j] /= num
		}

		x_[idx[i]] /= num

		for j := range a_ {

			if idx[i] == idx[j] {
				continue
			}

			per := a_[idx[j]][i] / a_[idx[i]][i]

			for k := i; k < len(a_); k++ {
				a_[idx[j]][k] -= per * a_[idx[i]][k]
			}

			x_[idx[j]] -= per * x_[idx[i]]
		}

	}

	for i, changed := range idx {
		x[changed] = x_[i]
	}

	return
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
