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

func backSubIdx(a Mat, b Vec, idx []int) (Vec, error) {
	a_ := NewMat(a.col, a.row)
	b_ := NewVec(b.row)

	for i := 0; i < a.col; i++ {
		a_.m[i] = a.m[idx[i]]
		b_.Write(i, b.At(idx[i]))
	}

	return BackSub(a_, b_)
}

func canSimultaneousEquSolve(a Mat, b Vec) bool {
	return a.IsSquareMat() && (a.col == b.row)
}

func forwardDelIdx(a Mat, b Vec, idx []int) error {
	for i := 0; i < a.col; i++ {

		if a.At(idx[i], i) == 0 {
			if i == a.col-1 {
				return ErrCannotSolve
			}

			tmp := make([]float64, 0)
			for j := i + 1; j < a.col; j++ {
				tmp = append(tmp, a.At(idx[j], i))
			}
			tmp2 := NewVecSet(tmp...)

			if maxnum, maxidx := max(tmp2); maxnum != 0 {
				idx[i], idx[maxidx+i+1] = idx[maxidx+i+1], idx[i]
			} else {
				return ErrCannotSolve
			}
		}

		for j := i + 1; j < a.col; j++ {
			per := a.At(idx[j], i) / a.At(idx[i], i)
			for k := i; k < a.col; k++ {
				a.Write(idx[j], k, a.At(idx[j], k)-per*a.At(idx[i], k))
			}

			b.Write(idx[j], b.At(idx[j])-per*b.At(idx[i]))
		}
	}
	return nil
}

func max(x Vec) (float64, int) {
	ans := x.At(0)
	idx := 0
	for i := 1; i < x.row; i++ {
		if math.Abs(x.At(i)) > math.Abs(ans) {
			ans = x.At(i)
			idx = i
		}
	}

	return ans, idx
}

func makeIdentityMat(n int) Mat {
	mat := NewMat(n, n)
	mat.Write(-1, -1, 1)
	return mat
}

func swapMatIdx(x Mat, idx []int) Mat {
	x_ := NewMat(x.col, x.row)

	for i, v := range idx {
		x_.m[i] = x.m[v]
	}
	return x_
}
