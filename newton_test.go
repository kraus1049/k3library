package k3library

import (
	"math"
	"testing"
)

type newtonTest struct {
	f                    func(float64) float64
	g                    func(float64) float64
	start, eps, expected float64
	ok                   bool
}

func TestNewton(t *testing.T) {
	// f g,start, eps, expected, ok
	testnewton := []newtonTest{
		{func(x float64) float64 { return math.Sin(x) },
			func(x float64) float64 { return math.Cos(x) },
			1, 1e-15, 0, true},
		{func(x float64) float64 { return math.Sin(x) },
			func(x float64) float64 { return math.Cos(x) },
			3, 1e-15, math.Pi, true},
	}

	for i := range testnewton {
		test := &testnewton[i]
		actual, ok := Newton(test.start, test.f, test.g, test.eps)

		if !Epsequal(actual, test.expected, test.eps) {
			t.Errorf("%v:actual = %v, expected = %v", i, actual, test.expected)
		} else if ok != test.ok {
			t.Errorf("%v:actual = %v,expected = %v", i, ok, test.ok)
		}

	}

}

func TestNewtonInvalidArgument(t *testing.T) {
	// f g,start, eps, expected, ok
	testnewtoninvalidargument := []newtonTest{
		{func(x float64) float64 { return math.Sin(x) },
			func(x float64) float64 { return math.Cos(x) },
			math.Pi / 2.0, 1e-15, 0, false},
		{func(x float64) float64 { return math.Tanh(x) },
			func(x float64) float64 { return 1 / (1 + math.Pow(x, 2)) },
			1.6, 1e-15, 0, false},
	}

	for i := range testnewtoninvalidargument {
		test := &testnewtoninvalidargument[i]
		_, ok := Newton(test.start, test.f, test.g, test.eps)

		if ok != test.ok {
			t.Errorf("%v:actual = %v,expected = %v", i, ok, test.ok)
		}
	}
}
