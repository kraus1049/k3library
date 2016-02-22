package k3library

import (
	"math"
	"testing"
)

type bisectTest struct {
	f                              func(float64) float64
	start, end, num, eps, expected float64
	ok                             bool
}

func TestBisect(t *testing.T) {

	// f ,start, end, num, eps, expected, ok
	var testbisect = []bisectTest{
		{func(x float64) float64 { return x - 1 },
			-3, 2, 0, 1e-15, 1, true},
		{func(x float64) float64 { return math.Sin(x) },
			-math.Pi / 2.0, math.Pi / 3.0, 0, 1e-15, 0, true},
		{func(x float64) float64 { return math.Pow(x, 2) },
			0, 20, 4, 1e-15, 2, true},
		{func(x float64) float64 { return math.Sin(x) },
			4, 3, 0, 1e-15, math.Pi, true},
	}

	for i := range testbisect {
		test := &testbisect[i]
		actual, ok := Bisect(test.start, test.end, test.f, test.num, test.eps)

		if !Epsequal(actual, test.expected, test.eps) {
			t.Errorf("%v:actual = %v, expected %v\n", i, actual, test.expected)
		} else if ok != test.ok {
			t.Errorf("%v:actual = %v, expected %v\n", i, ok, test.ok)
		}
	}
}

func TestBisectInvalidArgument(t *testing.T) {

	// f ,start, end, num, eps, expected, ok
	testbisectinvalidargument := []bisectTest{
		{func(x float64) float64 { return x - 1 },
			2, 3, 0, 1e-15, 0, false},
		{func(x float64) float64 { return 1 / (x - 5) },
			0, 10, 0, 1e-15, 0, false},
	}

	for i := range testbisectinvalidargument {
		test := &testbisectinvalidargument[i]
		_, ok := Bisect(test.start, test.end, test.f, test.num, test.eps)
		if ok != test.ok {
			t.Errorf("%v:actual = %v, expected %v\n", i, ok, test.ok)
		}
	}

}
