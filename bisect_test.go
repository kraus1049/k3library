package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"math"
	"testing"
)

type bisectTest struct {
	f                              func(float64) float64
	start, end, num, eps, expected float64
	err                            error
}

func TestBisect(t *testing.T) {

	// f ,start, end, num, eps, expected, err
	var testbisect = []bisectTest{
		{func(x float64) float64 { return x - 1 },
			-3, 2, 0, 1e-15, 1, nil},
		{func(x float64) float64 { return math.Sin(x) },
			-math.Pi / 2.0, math.Pi / 3.0, 0, 1e-15, 0, nil},
		{func(x float64) float64 { return math.Pow(x, 2) },
			0, 20, 4, 1e-15, 2, nil},
		{func(x float64) float64 { return math.Sin(x) },
			4, 3, 0, 1e-15, math.Pi, nil},
	}

	for i := range testbisect {
		test := &testbisect[i]
		actual, err := Bisect(test.start, test.end, test.f, test.num, test.eps)

		if !Epsequal(actual, test.expected, test.eps) {
			t.Errorf("%v:actual = %v, expected %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v:actual = %v, expected %v\n", i, err, test.err)
		}
	}
}

func TestBisectInvalidArgument(t *testing.T) {

	// f ,start, end, num, eps, expected, err
	testbisectinvalidargument := []bisectTest{
		{func(x float64) float64 { return x - 1 },
			2, 3, 0, 1e-15, 0,
			ErrInvalid},
		{func(x float64) float64 { return 1 / (x - 5) },
			0, 10, 0, 1e-15, 0,
			ErrInfiniteLoop},
	}

	for i := range testbisectinvalidargument {
		test := &testbisectinvalidargument[i]
		_, err := Bisect(test.start, test.end, test.f, test.num, test.eps)
		if err != test.err {
			t.Errorf("%v:actual = %v, expected %v\n", i, err, test.err)
		}
	}
}
