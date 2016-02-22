package k3library

import (
	"math"
	"testing"
)

type brentTest struct {
	f                              func(float64) float64
	start, end, num, eps, expected float64
	err                            error
}

func TestBrent(t *testing.T) {
	var testbrent = []brentTest{
		{func(x float64) float64 { return x - 1 },
			-3, 2, 0, 1e-15, 1, nil},
		{func(x float64) float64 { return math.Sin(x) },
			3, 4, 0, 1e-15, math.Pi, nil},
		{func(x float64) float64 { return math.Log(x) },
			0.5, 3, 1, 1e-15, math.E, nil},
	}

	for i := range testbrent {
		test := &testbrent[i]
		actual, err := Brent(test.start, test.end, test.f, test.num, test.eps)

		if !Epsequal(actual, test.expected, test.eps) {
			t.Errorf("%v:actual = %v, expected = %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v: actual = %v, expected %v\n", i, err, test.err)
		}
	}
}

func TestBrentInvalidArgument(t *testing.T) {
	var testbrentinvalidargument = []brentTest{
		{func(x float64) float64 { return x },
			1, 2, 0, 1e-15, 0,
			ErrInvalid},
	}

	for i := range testbrentinvalidargument {
		test := &testbrentinvalidargument[i]
		_, err := Brent(test.start, test.end, test.f, test.num, test.eps)

		if err != test.err {
			t.Errorf("%v:actual = %v,expected = %v\n", i, err, test.err)
		}
	}
}
