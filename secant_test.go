package k3library

import (
	"math"
	"testing"
)

type secantTest struct {
	f                    func(float64) float64
	start, end, num, eps float64
	expected             float64
	err                  error
}

func TestSecant(t *testing.T) {
	// f, start, end, num, eps, expected ,err
	testsecant := []secantTest{
		{func(x float64) float64 { return x },
			-1, 1, 0, 1e-15, 0, nil},
		{func(x float64) float64 { return math.Sin(x) },
			3, 4, 0, 1e-15, math.Pi, nil},
		{func(x float64) float64 { return math.Log(x) },
			5, 1, 1, 1e-15, math.E, nil},
	}

	for i := range testsecant {
		test := &testsecant[i]

		actual, err := Secant(test.start, test.end, test.f, test.num, test.eps)

		if !Epsequal(actual, test.expected, test.eps) {
			t.Errorf("actual = %v,expected = %v", actual, test.expected)
		} else if err != test.err {
			t.Errorf("actual = %v,expected = %v", err, test.err)
		}
	}
}

func testSecantinvalidargument(t *testing.T) {
	// f, start, end, num, eps, expected ,err
	testsecantinvalidargument := []secantTest{
		{func(x float64) float64 { return x },
			1, 1, 0, 1e-15, 0,
			ErrInvalid},
	}

	for i := range testsecantinvalidargument {
		test := &testsecantinvalidargument[i]

		_, err := Secant(test.start, test.end, test.f, test.num, test.eps)

		if err != test.err {
			t.Errorf("actual = %v,expected = %v", err, test.err)
		}
	}
}
