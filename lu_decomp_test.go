package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type luDecompTest struct {
	mat [][]float64
	err error
}

func TestLUDecomp(t *testing.T) {
	var testLUDecomp = []luDecompTest{
		{[][]float64{{1, 2}, {3, 4}},
			nil},
		{[][]float64{{3, 1, 4}, {1, 5, 9}, {2, 6, 5}},
			nil},
		{[][]float64{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}},
			nil},
		{[][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			nil},
		{[][]float64{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}},
			nil},
	}

	for i := range testLUDecomp {
		test := &testLUDecomp[i]
		l, u, idx, sgn, err := LUDecomp(test.mat)
		swap_mat := swapMatIdx(test.mat, idx)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		} else if pro, _ := MPro(l, u); !MatEpsEqual(pro, swap_mat, 1e-8) {
			t.Errorf("%v: l = %v , u = %v lu = %v, expected = %v\nsgn = %v, idx = %v\n", i, l, u, pro, swap_mat, sgn, idx)
		}
	}
}

func swapMatIdx(x [][]float64, idx []int) [][]float64 {
	x_ := make([][]float64, len(x))

	for i, v := range idx {
		x_[i] = x[v]
	}
	return x_
}
