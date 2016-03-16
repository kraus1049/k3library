package k3library

import (
	"testing"
)

type gauss_eliminationTest struct {
	a           [][]float64
	b, expected []float64
	err         error
}

func TestGauss_elimination(t *testing.T) {
	var testgauss_elimination = []gauss_eliminationTest{
		{[][]float64{{2, 1}, {1, -1}},
			[]float64{7, -1},
			[]float64{2, 3},
			nil},
		{[][]float64{{1, 1, 1}, {1, -1, 2}, {2, -3, 5}},
			[]float64{3, 2, 4},
			[]float64{1, 1, 1},
			nil},
		{[][]float64{{1, 1, 1}, {2, 2, -1}, {1, -1, 2}},
			[]float64{6, 3, 5},
			[]float64{1, 2, 3},
			nil},
		{[][]float64{{0, 1, 1, 1}, {1, -1, 0, -1}, {10, -5, -1, 0}, {3, 0, 2, 2}},
			[]float64{9, -5, -3, 17},
			[]float64{1, 2, 3, 4},
			nil},
	}

	for i := range testgauss_elimination {
		test := &testgauss_elimination[i]
		actual, err := Gauss_elimination(test.a, test.b)

		if !slice_epsequal(test.expected, actual, 1e-8) {
			t.Errorf("%v:actual = %v, expected = %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v:actual = %v,expected = %v\n", i, err, test.err)
		}
	}
}

func TestGauss_eliminationInvalidArgument(t *testing.T) {
	var testgauss_elimination_invalidargument = []gauss_eliminationTest{
		{[][]float64{{1, 2}, {3, 4, 5}},
			[]float64{1, 2, 3},
			[]float64{1, 2, 3},
			ErrInvalid},
		{[][]float64{{1, 2}, {3, 4}},
			[]float64{1, 2, 3},
			[]float64{1, 2, 3},
			ErrInvalid},
		{[][]float64{{1, 2, 3, 4}, {5, 6, 7}, {8, 9}, {10}},
			[]float64{1, 2, 3, 4},
			[]float64{1, 2, 3, 4},
			ErrInvalid},
	}

	for i := range testgauss_elimination_invalidargument {
		test := &testgauss_elimination_invalidargument[i]
		_, err := Gauss_elimination(test.a, test.b)

		if err != test.err {
			t.Errorf("%v:actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}

func TestGauss_eliminationCannotSolve(t *testing.T) {
	var testGauss_elimination_cannotsolve = []gauss_eliminationTest{
		{[][]float64{{1, 1}, {2, 2}},
			[]float64{1, 2},
			[]float64{1, 1},
			ErrCannotSolve},
		{[][]float64{{1, 1}, {1, 1}},
			[]float64{1, 2},
			[]float64{1, 1},
			ErrCannotSolve},
	}

	for i := range testGauss_elimination_cannotsolve {
		test := &testGauss_elimination_cannotsolve[i]
		_, err := Gauss_elimination(test.a, test.b)

		if err != test.err {
			t.Errorf("%v:actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}

func slice_epsequal(x, y []float64, eps float64) bool {
	for i := range x {
		if !Epsequal(x[i], y[i], eps) {
			return false
		}
	}

	return true
}
