package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type gauss_jordanTest struct {
	a           Mat
	b, expected Vec
	err         error
}

func TestGaussJordan(t *testing.T) {
	var testgauss_jordan = []gauss_jordanTest{
		{NewMatSet([][]float64{{2, 1}, {1, -1}}),
			NewVecSet([]float64{7, -1}),
			NewVecSet([]float64{2, 3}),
			nil},
		{NewMatSet([][]float64{{1, 1, 1}, {1, -1, 2}, {2, -3, 5}}),
			NewVecSet([]float64{3, 2, 4}),
			NewVecSet([]float64{1, 1, 1}),
			nil},
		{NewMatSet([][]float64{{1, 1, 1}, {2, 2, -1}, {1, -1, 2}}),
			NewVecSet([]float64{6, 3, 5}),
			NewVecSet([]float64{1, 2, 3}),
			nil},
		{NewMatSet([][]float64{{0, 1, 1, 1}, {1, -1, 0, -1}, {10, -5, -1, 0}, {3, 0, 2, 2}}),
			NewVecSet([]float64{9, -5, -3, 17}),
			NewVecSet([]float64{1, 2, 3, 4}),
			nil},
	}

	for i := range testgauss_jordan {
		test := &testgauss_jordan[i]
		actual, err := GaussJordan(test.a, test.b)

		if !VecEpsEqual(test.expected, actual, 1e-8) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v:actual = %v,expected = %v\n", i, err, test.err)
		}
	}
}

func TestGaussJordanInvalidArgument(t *testing.T) {
	var testgauss_jordan_invalidargument = []gauss_jordanTest{
		{NewMatSet([][]float64{{1, 2}, {3, 4, 5}}),
			NewVecSet([]float64{1, 2, 3}),
			NewVecSet([]float64{1, 2, 3}),
			ErrInvalid},
		{NewMatSet([][]float64{{1, 2}, {3, 4}}),
			NewVecSet([]float64{1, 2, 3}),
			NewVecSet([]float64{1, 2, 3}),
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
		{NewMatSet([][]float64{{1, 1}, {2, 2}}),
			NewVecSet([]float64{1, 2}),
			NewVecSet([]float64{1, 1}),
			ErrCannotSolve},
		{NewMatSet([][]float64{{1, 1}, {1, 1}}),
			NewVecSet([]float64{1, 2}),
			NewVecSet([]float64{1, 1}),
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
