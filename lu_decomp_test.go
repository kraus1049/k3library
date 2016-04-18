package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type luDecompTest struct {
	mat Mat
	err error
}

func TestLUDecomp(t *testing.T) {
	var testLUDecomp = []luDecompTest{
		{NewMatSet([][]float64{{1, 2}, {3, 4}}),
			nil},
		{NewMatSet([][]float64{{3, 1, 4}, {1, 5, 9}, {2, 6, 5}}),
			nil},
		{NewMatSet([][]float64{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}),
			nil},
		{NewMatSet([][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 0}}),
			nil},
		{NewMatSet([][]float64{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}}),
			nil},
		{NewMatSet([][]float64{{0, 1, 1, 1}, {1, -1, 0, -1}, {10, -5, -1, 0}, {3, 0, 2, 2}}),
			nil},
	}

	for i := range testLUDecomp {
		test := &testLUDecomp[i]
		l, u, idx, sgn, err := test.mat.LUDecomp()
		swap_mat := swapMatIdx(test.mat, idx)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		} else if pro, _ := Pro(l, u); !MatEpsEqual(pro.(Mat), swap_mat, 1e-8) {
			t.Errorf("%v: l = %v , u = %v lu = %v, expected = %v\nsgn = %v, idx = %v\n", i, l, u, pro, swap_mat, sgn, idx)
		}
	}
}

func swapMatIdx(x Mat, idx []int) Mat {
	x_ := NewMat(x.Col(), x.Row())

	for i, v := range idx {
		x_.M()[i] = x.M()[v]
	}
	return x_
}
