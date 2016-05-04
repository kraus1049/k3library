package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type leastSquaresTest struct {
	data     Mat
	degree   int
	expected Vec
	err      error
}

func TestLeastSquares(t *testing.T) {
	var testLeastSquares = []leastSquaresTest{
		{NewMatSet([][]float64{{2, 3}, {4, 7}, {9, 11}}),
			1,
			NewVecSet(21.0/13, 14.0/13),
			nil},
		{NewMatSet([][]float64{{2, 3}, {4, 7}, {9, 11}}),
			2,
			NewVecSet(-2.3714285714283294, 3.0285714285714054, -0.1714285714285699),
			nil},
	}

	for i := range testLeastSquares {
		test := &testLeastSquares[i]
		actual, err := LeastSquares(test.data, test.degree)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v", i, err, test.err)
		} else if !VecEpsEqual(actual, test.expected, 1e-5) {
			t.Errorf("%v: actual = %v, expected = %v", i, actual, test.expected)
		}
	}
}
