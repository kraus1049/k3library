package k3library

import (
	"math"
	"testing"
)

type regula_falsaTest struct {
	f                    func(float64) float64
	start, end, num, eps float64
	expected             float64
	ok                   bool
}

func TestRegula_falsa(t *testing.T) {
	// f, start, end, num, eps, expected ,ok
	testregula_falsa := []regula_falsaTest{
		{func(x float64) float64 { return x },
			-1, 1, 0, 1e-15, 0, true},
		{func(x float64) float64 { return math.Sin(x) },
			3, 4, 0, 1e-15, math.Pi, true},
		{func(x float64) float64 { return math.Log(x) },
			5, 1, 1, 1e-15, math.E, true},
	}

	for i := range testregula_falsa {
		test := &testregula_falsa[i]

		actual, ok := Regula_falsa(test.start, test.end, test.f, test.num, test.eps)

		if !Epsequal(actual, test.expected, test.eps) {
			t.Errorf("actual = %v,expected = %v", actual, test.expected)
		} else if ok != test.ok {
			t.Errorf("actual = %v,expected = %v", ok, test.ok)
		}
	}
}

func TestRegula_falsaInvalidArgument(t *testing.T) {
	// f, start, end, num, eps, expected ,ok
	testregula_falsainvalidargument := []regula_falsaTest{
		{func(x float64) float64 { return x },
			1, 1, 0, 1e-15, 0, false},
		{func(x float64) float64 { return math.Pow(x, 2) - 1 },
			2, 3, 0, 1e-15, 0, false},
	}

	for i := range testregula_falsainvalidargument {
		test := &testregula_falsainvalidargument[i]

		_, ok := Regula_falsa(test.start, test.end, test.f, test.num, test.eps)

		if ok != test.ok {
			t.Errorf("actual = %v,expected = %v", ok, test.ok)
		}
	}
}
