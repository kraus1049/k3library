package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"math"
	"testing"
)

type rectangleTest struct {
	f                   func(float64) float64
	a, b, eps, expected float64
	err                 error
}

func TestRectangle(t *testing.T) {
	var testRectangle = []rectangleTest{
		{func(x float64) float64 { return x },
			0, 10, 1e-5, 50,
			nil},
		{func(x float64) float64 { return 10 },
			5, 10, 1e-5, 50,
			nil},
		{func(x float64) float64 { return math.Sin(x) },
			math.Pi / 2, math.Pi + math.Pi/2, 1e-5, 0,
			nil},
	}

	for i := range testRectangle {
		test := &testRectangle[i]
		actual, err := Rectangle(Lift(test.f), test.a, test.b, test.eps)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		} else if !EpsEqual(actual, test.expected, test.eps) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		}
	}

}

func TestRectangleInvalidArgument(t *testing.T) {
	var testRectangleInvalidArgument = []rectangleTest{
		{func(x float64) float64 { return x },
			10, 0, 1e-5, 50,
			ErrInvalid},
		{func(x float64) float64 { return 10 },
			5, 5, 1e-5, 50,
			ErrInvalid},
	}

	for i := range testRectangleInvalidArgument {
		test := &testRectangleInvalidArgument[i]
		_, err := Rectangle(Lift(test.f), test.a, test.b, test.eps)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}

func BenchmarkRectangle(b *testing.B) {
	f := func(x float64) float64 {
		return math.Sin(x)
	}

	a := 0.0
	c := math.Pi
	eps := 1e-5
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Rectangle(Lift(f), a, c, eps)
	}

}
