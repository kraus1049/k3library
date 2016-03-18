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
	}

	for i := range testLUDecomp {
		test := &testLUDecomp[i]
		l, u, err := LUDecomp(test.mat)

		if pro, _ := MPro(l, u); !MatEpsEqual(pro, test.mat, 1e-8) {
			t.Errorf("%v: l = %v , u = %v lu = %v, expected = %v", i, l, u, pro, test.mat)
		} else if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}
