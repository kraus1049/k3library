package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type forwardSubTest struct {
	mat      [][]float64
	b        []float64
	expected []float64
	err      error
}

func TestForwardSub(t *testing.T) {
	var testForwardSub = []forwardSubTest{
		{[][]float64{{1, 0}, {1, 1}},
			[]float64{1, 2},
			[]float64{1, 1},
			nil,
		},
		{[][]float64{{1, 0, 0}, {2, 3, 0}, {4, 5, 6}},
			[]float64{1, 5, 15},
			[]float64{1, 1, 1},
			nil,
		},
		{[][]float64{{1, 123, 321}, {2, 3, 456}, {4, 5, 6}},
			[]float64{1, 5, 15},
			[]float64{1, 1, 1},
			nil,
		},
	}

	for i := range testForwardSub {
		test := &testForwardSub[i]
		actual, err := ForwardSub(test.mat, test.b)

		if !SliceEpsEqual(actual, test.expected, 1e-8) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}

func TestForwardSubCannotSolve(t *testing.T) {
	var testForwardSubCannotSolve = []forwardSubTest{
		{[][]float64{{1, 0, 0}, {2, 0, 0}, {4, 5, 6}},
			[]float64{1, 5, 15},
			[]float64{1, 1, 1},
			ErrCannotSolve,
		},
	}

	for i := range testForwardSubCannotSolve {
		test := &testForwardSubCannotSolve[i]
		_, err := ForwardSub(test.mat, test.b)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}
