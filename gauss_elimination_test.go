package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type gauss_eliminationTest struct {
	a           Mat
	b, expected Vec
	err         error
}

func TestGaussElimination(t *testing.T) {
	var testgauss_elimination = []gauss_eliminationTest{
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

	for i := range testgauss_elimination {
		test := &testgauss_elimination[i]
		actual, err := GaussElimination(test.a, test.b)

		if !VecEpsEqual(test.expected, actual, 1e-8) {
			t.Errorf("%v:actual = %v, expected = %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v:actual = %v,expected = %v\n", i, err, test.err)
		}
	}
}

func TestGaussEliminationInvalidArgument(t *testing.T) {
	var testgauss_elimination_invalidargument = []gauss_eliminationTest{
		{NewMatSet([][]float64{{1, 2}, {3, 4, 5}}),
			NewVecSet([]float64{1, 2, 3}),
			NewVecSet([]float64{1, 2, 3}),
			ErrInvalid},
		{NewMatSet([][]float64{{1, 2}, {3, 4}}),
			NewVecSet([]float64{1, 2, 3}),
			NewVecSet([]float64{1, 2, 3}),
			ErrInvalid},
	}

	for i := range testgauss_elimination_invalidargument {
		test := &testgauss_elimination_invalidargument[i]
		_, err := GaussElimination(test.a, test.b)

		if err != test.err {
			t.Errorf("%v:actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}

func TestGaussEliminationCannotSolve(t *testing.T) {
	var testGaussElimination_cannotsolve = []gauss_eliminationTest{
		{NewMatSet([][]float64{{1, 1}, {2, 2}}),
			NewVecSet([]float64{1, 2}),
			NewVecSet([]float64{1, 1}),
			ErrCannotSolve},
		{NewMatSet([][]float64{{1, 1}, {1, 1}}),
			NewVecSet([]float64{1, 2}),
			NewVecSet([]float64{1, 1}),
			ErrCannotSolve},
	}

	for i := range testGaussElimination_cannotsolve {
		test := &testGaussElimination_cannotsolve[i]
		_, err := GaussElimination(test.a, test.b)

		if err != test.err {
			t.Errorf("%v:actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}
