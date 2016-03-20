package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type vTest struct {
	a, b     []float64
	expected []float64
	err      error
}

type mProTest struct {
	a, b     [][]float64
	expected [][]float64
	err      error
}

func TestVSub(t *testing.T) {
	var testVSub = []vTest{
		{[]float64{1, 2},
			[]float64{1, 1},
			[]float64{0, 1},
			nil,
		},
	}

	for i := range testVSub {
		test := &testVSub[i]
		actual, err := VSub(test.a, test.b)

		if !SliceEpsEqual(actual, test.expected, 1e-8) {
			t.Errorf("%v: actual = %v expected = %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v: actual = %v expected = %v\n", i, err, test.err)
		}
	}
}

func TestMPro(t *testing.T) {
	var testMPro = []mProTest{
		{[][]float64{{1, 1}, {1, 1}},
			[][]float64{{1, 1}, {1, 1}},
			[][]float64{{2, 2}, {2, 2}},
			nil,
		},
		{[][]float64{{1, 2}, {3, 4}},
			[][]float64{{5, 6}, {7, 8}},
			[][]float64{{19, 22}, {43, 50}},
			nil,
		},
		{[][]float64{{1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1}, {1, 1}, {1, 1}},
			[][]float64{{3, 3}, {3, 3}},
			nil,
		},
	}

	for i := range testMPro {
		test := &testMPro[i]
		actual, err := MPro(test.a, test.b)

		if !MatEpsEqual(actual, test.expected, 1e-8) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}

func TestMProInvalid(t *testing.T) {
	var testMProInvalid = []mProTest{
		{[][]float64{{1, 1}},
			[][]float64{{1, 1}},
			[][]float64{{1, 1}, {1, 1}},
			ErrInvalid,
		},
		{[][]float64{{1, 1}},
			[][]float64{{1, 1}},
			[][]float64{{1, 1}, {1, 1}},
			ErrInvalid,
		},
	}

	for i := range testMProInvalid {
		test := &testMProInvalid[i]
		_, err := MPro(test.a, test.b)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}
