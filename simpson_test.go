package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"math"
	"testing"
)

type simpsonTest struct {
	f                   func(float64) float64
	a, b, eps, expected float64
	err                 error
}

func TestSimpson(t *testing.T) {
	var testSimpson = []simpsonTest{
		{func(x float64) float64 { return x },
			0, 10, 1e-6, 50,
			nil},
		{func(x float64) float64 { return 10 },
			5, 10, 1e-6, 50,
			nil},
		{func(x float64) float64 { return math.Sin(x) },
			0, math.Pi / 2, 1e-6, 1,
			nil},
	}

	for i := range testSimpson {
		test := &testSimpson[i]
		actual, err := Simpson(Lift(test.f), test.a, test.b, test.eps)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		} else if !EpsEqual(actual, test.expected, test.eps) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		}
	}
}

func TestSimpsonInvalidArgument(t *testing.T) {
	var testSimpsonInvalidArgument = []simpsonTest{
		{func(x float64) float64 { return x },
			10, 0, 1e-6, 50,
			ErrInvalid},
		{func(x float64) float64 { return 10 },
			5, 5, 1e-5, 50,
			ErrInvalid},
	}

	for i := range testSimpsonInvalidArgument {
		test := &testSimpsonInvalidArgument[i]
		_, err := Simpson(Lift(test.f), test.a, test.b, test.eps)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}

}
