package k3library

import (
	"math"
	"testing"
)

type zbracTest struct {
	f          func(float64) float64
	start, end float64
	expected   float64
	ok         bool
}

func TestZbrac(t *testing.T) {

	// f start, end expected  ok
	testzbrac := []zbracTest{
		{func(x float64) float64 { return x },
			-1, 1, 0, true},
		{func(x float64) float64 { return math.Sin(x) },
			1, 1.5, -math.Pi / 2.0, true},
		{func(x float64) float64 { return math.Pow(x, 3) - 3*math.Pow(x, 2) + 3 },
			0, 1, 2, true},
	}

	for i := range testzbrac {
		test := &testzbrac[i]
		actual_1, actual_2, ok := Zbrac(test.start, test.end, test.f)

		if test.expected < actual_1 || actual_2 < test.expected {
			t.Errorf("%v:I want actual_1(%v) < expected(%v) < actual_2(%v)\n", i, actual_1, test.expected, actual_2)
		} else if ok != test.ok {
			t.Errorf("%v:actual = %v,expected = %v\n", ok, test.ok)
		}
	}
}

func TestZbracBadInitialRange(t *testing.T) {

	// f start, end expected  ok
	testzbracbadinitialrange := []zbracTest{
		{func(x float64) float64 { return math.Tanh(x) },
	1,1,0,false},
	}

	for i := range testzbracbadinitialrange {
		test := &testzbracbadinitialrange[i]
		_, _, ok := Zbrac(test.start, test.end, test.f)

		if ok != test.ok {
			t.Errorf("%v:actual = %v,expected = %v\n", ok, test.ok)
		}
	}

}
