package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type gauss_jordanTest struct {
	a           [][]float64
	b, expected []float64
	err         error
}

func TestGaussJordan(t *testing.T) {
	var testgauss_jordan = []gauss_jordanTest{
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

	for i := range testgauss_jordan {
		test := &testgauss_jordan[i]
		actual, err := GaussJordan(test.a, test.b)

		if !slice_epsequal(test.expected, actual, 1e-8) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v:actual = %v,expected = %v\n", i, err, test.err)
		}
	}
}

func TestGaussJordanInvalidArgument(t *testing.T) {
	var testgauss_jordan_invalidargument = []gauss_jordanTest{
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

	for i := range testgauss_jordan_invalidargument {
		test := &testgauss_jordan_invalidargument[i]
		_, err := GaussJordan(test.a, test.b)

		if err != test.err {
			t.Errorf("%v:actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}

func TestGaussJordanCannotSolve(t *testing.T) {
	var testgauss_jordan_cannotsolve = []gauss_jordanTest{
		{[][]float64{{1, 1}, {2, 2}},
			[]float64{1, 2},
			[]float64{1, 1},
			ErrCannotSolve},
		{[][]float64{{1, 1}, {1, 1}},
			[]float64{1, 2},
			[]float64{1, 1},
			ErrCannotSolve},
	}

	for i := range testgauss_jordan_cannotsolve {
		test := &testgauss_jordan_cannotsolve[i]
		_, err := GaussJordan(test.a, test.b)

		if err != test.err {
			t.Errorf("%v:actual = %v, expected = %v\n", i, err, test.err)
		}
	}

}
