package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type backSubTest struct {
	mat      Mat
	b        Vec
	expected Vec
	err      error
}

func TestBackSub(t *testing.T) {
	var testBackSub = []backSubTest{
		{NewMatSet([][]float64{{1, 1}, {0, 1}}),
			NewVecSet([]float64{2, 1}),
			NewVecSet([]float64{1, 1}),
			nil,
		},
		{NewMatSet([][]float64{{4, 5, 6}, {0, 2, 3}, {0, 0, 1}}),
			NewVecSet([]float64{15, 5, 1}),
			NewVecSet([]float64{1, 1, 1}),
			nil,
		},
		{NewMatSet([][]float64{{4, 5, 6}, {123, 2, 3}, {345, 987, 1}}),
			NewVecSet([]float64{15, 5, 1}),
			NewVecSet([]float64{1, 1, 1}),
			nil,
		},
	}

	for i := range testBackSub {
		test := &testBackSub[i]
		actual, err := BackSub(test.mat, test.b)

		if !VecEpsEqual(actual, test.expected, 1e-8) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}

func TestBackSubCannotSolve(t *testing.T) {
	var testBackSubCannotSolve = []backSubTest{
		{NewMatSet([][]float64{{0, 1}, {0, 1}}),
			NewVecSet([]float64{2, 1}),
			NewVecSet([]float64{1, 1}),
			ErrCannotSolve,
		},
		{NewMatSet([][]float64{{4, 5, 6}, {0, 0, 3}, {0, 0, 1}}),
			NewVecSet([]float64{15, 5, 1}),
			NewVecSet([]float64{1, 1, 1}),
			ErrCannotSolve,
		},
	}

	for i := range testBackSubCannotSolve {
		test := &testBackSubCannotSolve[i]
		_, err := BackSub(test.mat, test.b)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}
