package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"math"
	"testing"
)

type trapezoidTest struct {
	f                   func(float64) float64
	a, b, eps, expected float64
	err                 error
}

func TestTrapezoid(t *testing.T) {
	var testTrapezoid = []trapezoidTest{
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

	for i := range testTrapezoid {
		test := &testTrapezoid[i]
		actual, err := Trapezoid(test.f, test.a, test.b, test.eps)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		} else if !EpsEqual(actual, test.expected, test.eps) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		}
	}
}

func TestTrapezoidInvalidArgument(t *testing.T) {
	var testTrapezoidInvalidArgument = []trapezoidTest{
		{func(x float64) float64 { return x },
			10, 0, 1e-6, 50,
			ErrInvalid},
		{func(x float64) float64 { return 10 },
			5, 5, 1e-5, 50,
			ErrInvalid},
	}

	for i := range testTrapezoidInvalidArgument {
		test := &testTrapezoidInvalidArgument[i]
		_, err := Trapezoid(test.f, test.a, test.b, test.eps)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}

}
