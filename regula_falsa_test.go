package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"math"
	"testing"
)

type regula_falsaTest struct {
	f                    func(float64) float64
	start, end, num, eps float64
	expected             float64
	err                  error
}

func TestRegulaFalsa(t *testing.T) {
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

		actual, err := RegulaFalsa(test.start, test.end, test.f, test.num, test.eps)

		if !EpsEqual(actual, test.expected, test.eps) {
			t.Errorf("actual = %v,expected = %v", actual, test.expected)
		} else if err != test.err {
			t.Errorf("actual = %v,expected = %v", err, test.err)
		}
	}
}

func TestRegulaFalsaInvalidArgument(t *testing.T) {
	// f, start, end, num, eps, expected ,err
	testregula_falsainvalidargument := []regula_falsaTest{
		{func(x float64) float64 { return x },
			1, 1, 0, 1e-15, 0,
			ErrInvalid},
		{func(x float64) float64 { return math.Pow(x, 2) - 1 },
			2, 3, 0, 1e-15, 0,
			ErrInvalid},
	}

	for i := range testregula_falsainvalidargument {
		test := &testregula_falsainvalidargument[i]

		_, err := RegulaFalsa(test.start, test.end, test.f, test.num, test.eps)

		if err != test.err {
			t.Errorf("actual = %v,expected = %v", err, test.err)
		}
	}
}
