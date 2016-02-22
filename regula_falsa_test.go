package k3library

import (
	"errors"
	"math"
	"testing"
)

type regula_falsaTest struct {
	f                    func(float64) float64
	start, end, num, eps float64
	expected             float64
	err                   error
}

func TestRegula_falsa(t *testing.T) {
	// f, start, end, num, eps, expected ,err
	testregula_falsa := []regula_falsaTest{
		{func(x float64) float64 { return x },
			-1, 1, 0, 1e-15, 0, nil},
		{func(x float64) float64 { return math.Sin(x) },
			3, 4, 0, 1e-15, math.Pi, nil},
		{func(x float64) float64 { return math.Log(x) },
			5, 1, 1, 1e-15, math.E, nil},
	}

	for i := range testregula_falsa {
		test := &testregula_falsa[i]

		actual, err := Regula_falsa(test.start, test.end, test.f, test.num, test.eps)

		if !Epsequal(actual, test.expected, test.eps) {
			t.Errorf("actual = %v,expected = %v", actual, test.expected)
		} else if err != test.err {
			t.Errorf("actual = %v,expected = %v", err, test.err)
		}
	}
}

func TestRegula_falsaInvalidArgument(t *testing.T) {
	// f, start, end, num, eps, expected ,err
	testregula_falsainvalidargument := []regula_falsaTest{
		{func(x float64) float64 { return x },
			1, 1, 0, 1e-15, 0,
			errors.New("Regula_falsa:Invalid argument")},
		{func(x float64) float64 { return math.Pow(x, 2) - 1 },
			2, 3, 0, 1e-15, 0,
			errors.New("Regula_falsa:Invalid argument")},
	}

	for i := range testregula_falsainvalidargument {
		test := &testregula_falsainvalidargument[i]

		_, err := Regula_falsa(test.start, test.end, test.f, test.num, test.eps)

		if err.Error() != test.err.Error() {
			t.Errorf("actual = %v,expected = %v", err, test.err)
		}
	}
}
