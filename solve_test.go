package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type solveTest struct {
	mat      Mat
	b        Vec
	expected Vec
	err      error
}

func TestSolve(t *testing.T) {
	var testSolve = []solveTest{
		{NewMatSet([][]float64{{1, 0}, {0, 1}}),
			NewVecSet([]float64{1, 1}),
			NewVecSet([]float64{1, 1}),
			nil,
		},
		{NewMatSet([][]float64{{2, 1}, {1, -1}}),
			NewVecSet([]float64{7, -1}),
			NewVecSet([]float64{2, 3}),
			nil},
		{NewMatSet([][]float64{{1, 1, 1}, {1, -1, 2}, {2, -3, 5}}),
			NewVecSet([]float64{3, 2, 4}),
			NewVecSet([]float64{1, 1, 1}),
			nil},
		{NewMatSet([][]float64{{0, 1, 1, 1}, {1, -1, 0, -1}, {10, -5, -1, 0}, {3, 0, 2, 2}}),
			NewVecSet([]float64{9, -5, -3, 17}),
			NewVecSet([]float64{1, 2, 3, 4}),
			nil},
	}

	for i := range testSolve {
		test := &testSolve[i]

		l, u, idx, _, _ := test.mat.LUDecomp()
		actual, err := Solve(l, u, test.b, idx)

		if !VecEpsEqual(actual, test.expected, 1e-8) {
			t.Errorf("%v:actual = %v, expected = %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v: actual = %v, expected =%v\n", i, err, test.err)
		}
	}

}

func TestProveSolve(t *testing.T) {
	var testProveSolve = []solveTest{
		{NewMatSet([][]float64{{1, 0}, {0, 1}}),
			NewVecSet([]float64{1, 1}),
			NewVecSet([]float64{1, 1}),
			nil,
		},
		{NewMatSet([][]float64{{2, 1}, {1, -1}}),
			NewVecSet([]float64{7, -1}),
			NewVecSet([]float64{2, 3}),
			nil},
		{NewMatSet([][]float64{{1, 1, 1}, {1, -1, 2}, {2, -3, 5}}),
			NewVecSet([]float64{3, 2, 4}),
			NewVecSet([]float64{1, 1, 1}),
			nil},
		{NewMatSet([][]float64{{0, 1, 1, 1}, {1, -1, 0, -1}, {10, -5, -1, 0}, {3, 0, 2, 2}}),
			NewVecSet([]float64{9, -5, -3, 17}),
			NewVecSet([]float64{1, 2, 3, 4}),
			nil},
		{NewMatSet([][]float64{{9.876543219876543, 7.654321098765432},
			{1.35791357913689, 4.567891534534534534573}}),
			NewVecSet([]float64{30.14784272534673, 12.3912433503139}),
			NewVecSet([]float64{1.234567891234568, 2.345678912345679}),
			nil},
	}

	for i := range testProveSolve {
		test := &testProveSolve[i]
		actual, err := ProveSolve(test.mat, test.b, 1e-15)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		} else if !VecEpsEqual(actual, test.expected, 1e-10) {
			t.Errorf("%v: actual = %v,expected = %v\n", i, actual, test.expected)
		}
	}
}
