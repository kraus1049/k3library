package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type epsequalTest struct {
	a, b, eps float64
	expected  bool
}

type VecEpsEqualTest struct {
	a, b     Vec
	eps      float64
	expected bool
}

type MatEpsEqualTest struct {
	a, b     Mat
	eps      float64
	expected bool
}

func TestEpsEqual(t *testing.T) {

	// a,b,eps expected
	testepsequall := []epsequalTest{
		{1, 2, 3, true},
		{1, 2, 0, false},
		{1e-1, 11e-2, 1e-1, true},
		{1e-1, 11e-2, 1e-2, true},
		{1e-1, 11e-2, 1e-3, false},
	}

	for i := range testepsequall {
		test := &testepsequall[i]
		actual := EpsEqual(test.a, test.b, test.eps)

		if actual != test.expected {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		}
	}
}

func TestVecEpsEqual(t *testing.T) {
	var testVecEpsEqual = []VecEpsEqualTest{
		{NewVecSet([]float64{1, 1, 1}),
			NewVecSet([]float64{2, 2, 2}),
			3, true},
		{NewVecSet([]float64{1, 1, 1, 1}),
			NewVecSet([]float64{2, 2, 2, 2}),
			0, false},
		{NewVecSet([]float64{1, 1, 1, 1, 1}),
			NewVecSet([]float64{1, 1, 1}),
			0, false},
	}

	for i := range testVecEpsEqual {
		test := &testVecEpsEqual[i]
		actual := VecEpsEqual(test.a, test.b, test.eps)

		if actual != test.expected {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		}
	}

}

func TestMatEpsEqual(t *testing.T) {
	var testMatEpsEqual = []MatEpsEqualTest{
		{NewMatSet([][]float64{{1, 1}, {1, 1}}),
			NewMatSet([][]float64{{2, 2}, {2, 2}}),
			3, true},
		{NewMatSet([][]float64{{1, 1, 1}, {1, 1, 1}}),
			NewMatSet([][]float64{{2, 2, 2}, {2, 2, 2}}),
			0, false},
		{NewMatSet([][]float64{{1, 1, 1, 1}, {1, 1, 1, 1}}),
			NewMatSet([][]float64{{0, 0, 0}, {0, 0, 0}}),
			0, false},
	}

	for i := range testMatEpsEqual {
		test := &testMatEpsEqual[i]
		actual := MatEpsEqual(test.a, test.b, test.eps)

		if actual != test.expected {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		}
	}

}
