package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"math"
	"testing"
)

type newtonTest struct {
	f                    func(float64) float64
	g                    func(float64) float64
	start, eps, expected float64
	err                  error
}

func TestNewton(t *testing.T) {
	// f g,start, eps, expected, err
	testnewton := []newtonTest{
		{func(x float64) float64 { return math.Sin(x) },
			func(x float64) float64 { return math.Cos(x) },
			1, 1e-15, 0, nil},
		{func(x float64) float64 { return math.Sin(x) },
			func(x float64) float64 { return math.Cos(x) },
			3, 1e-15, math.Pi, nil},
	}

	for i := range testnewton {
		test := &testnewton[i]
		actual, err := Newton(test.start, test.f, test.g, test.eps)

		if !Epsequal(actual, test.expected, test.eps) {
			t.Errorf("%v:actual = %v, expected = %v", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v:actual = %v,expected = %v", i, err, test.err)
		}

	}

}

func TestNewtonInvalidArgument(t *testing.T) {
	// f g,start, eps, expected, err
	testnewtoninvalidargument := []newtonTest{
		{func(x float64) float64 { return math.Sin(x) },
			func(x float64) float64 { return math.Cos(x) },
			math.Pi / 2.0, 1e-15, 0,
			ErrInvalid},
		{func(x float64) float64 { return math.Tanh(x) },
			func(x float64) float64 { return 1 / (1 + math.Pow(x, 2)) },
			1.6, 1e-15, 0,
			ErrInvalid},
	}

	for i := range testnewtoninvalidargument {
		test := &testnewtoninvalidargument[i]
		_, err := Newton(test.start, test.f, test.g, test.eps)

		if err != test.err {
			t.Errorf("%v:actual = %v,expected = %v", i, err, test.err)
		}
	}
}
