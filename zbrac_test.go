package k3library

import (
	"errors"
	"math"
	"testing"
)

type zbracTest struct {
	f          func(float64) float64
	start, end float64
	expected   float64
	err         error
}

func TestZbrac(t *testing.T) {

	// f start, end expected  err
	testzbrac := []zbracTest{
		{func(x float64) float64 { return x },
			-1, 1, 0, nil},
		{func(x float64) float64 { return math.Sin(x) },
			1, 1.5, -math.Pi / 2.0, nil},
		{func(x float64) float64 { return math.Pow(x, 3) - 3*math.Pow(x, 2) + 3 },
			0, 1, 2, nil},
	}

	for i := range testzbrac {
		test := &testzbrac[i]
		actual_1, actual_2, err := Zbrac(test.start, test.end, test.f)

		if test.expected < actual_1 || actual_2 < test.expected {
			t.Errorf("%v:I want actual_1(%v) < expected(%v) < actual_2(%v)\n", i, actual_1, test.expected, actual_2)
		} else if err != test.err {
			t.Errorf("%v:actual = %v,expected = %v\n", err, test.err)
		}
	}
}

func TestZbracBadInitialRange(t *testing.T) {

	// f start, end expected  err
	testzbracbadinitialrange := []zbracTest{
		{func(x float64) float64 { return math.Tanh(x) },
	1,1,0,
	errors.New("Zbrac:Invalid argument")},
	}

	for i := range testzbracbadinitialrange {
		test := &testzbracbadinitialrange[i]
		_, _, err := Zbrac(test.start, test.end, test.f)

		if err.Error() != test.err.Error() {
			t.Errorf("%v:actual = %v,expected = %v\n", err, test.err)
		}
	}

}
