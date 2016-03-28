package k3library

import (
	"math"
	"reflect"
)

func allTypeEqual(a ...interface{}) bool {
	r := reflect.ValueOf(a[0])
	if r.IsValid() {

		for i := 1; i < len(a); i++ {
			rr := reflect.ValueOf(a[i])

			if rr.IsValid() {
				if r.Kind() != rr.Kind() {
					return false
				}
			}
		}
	}

	return true
}

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

func canSimultaneousEquSolve(a [][]float64, b []float64) bool {
	return IsSquareMat(a) && (len(a) == len(b))
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

func makeMat(col, row int) [][]float64 {
	mat := make([][]float64, col)
	for i := range mat {
		mat[i] = make([]float64, row)
	}
	return mat
}

func makeIdentityMat(n int) [][]float64 {
	mat := makeMat(n, n)
	for i := range mat {
		mat[i][i] = 1
	}
	return mat
}

func swapMatIdx(x [][]float64, idx []int) [][]float64 {
	x_ := make([][]float64, len(x))

	for i, v := range idx {
		x_[i] = x[v]
	}
	return x_
}
