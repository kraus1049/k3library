package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type solveTest struct {
	mat      [][]float64
	b        []float64
	expected []float64
	err      error
}

func TestSolve(t *testing.T) {
	var testSolve = []solveTest{
		{[][]float64{{1, 0}, {0, 1}},
			[]float64{1, 1},
			[]float64{1, 1},
			nil,
		},
		{[][]float64{{2, 1}, {1, -1}},
			[]float64{7, -1},
			[]float64{2, 3},
			nil},
		{[][]float64{{1, 1, 1}, {1, -1, 2}, {2, -3, 5}},
			[]float64{3, 2, 4},
			[]float64{1, 1, 1},
			nil},
		{[][]float64{{0, 1, 1, 1}, {1, -1, 0, -1}, {10, -5, -1, 0}, {3, 0, 2, 2}},
			[]float64{9, -5, -3, 17},
			[]float64{1, 2, 3, 4},
			nil},
	}

	for i := range testSolve {
		test := &testSolve[i]

		l, u, idx, _, _ := LUDecomp(test.mat)
		actual, err := Solve(l, u, test.b, idx)

		if !SliceEpsEqual(actual, test.expected, 1e-8) {
			t.Errorf("%v:actual = %v, expected = %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v: actual = %v, expected =%v\n", i, err, test.err)
		}
	}

}
