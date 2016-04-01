package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"math"
	"testing"
)

type eulerTest struct {
	f                 func(float64, float64) (float64, error)
	x_i, y_i, to, eps float64
	expected          func(float64) float64
	err               error
}

func TestEuler(t *testing.T) {
	var testEuler = []eulerTest{
		{func(x, y float64) (float64, error) { return 2 * x, nil },
			0, 0, 1, 1e-4,
			func(x float64) float64 { return math.Pow(x, 2) },
			nil},
		{func(x, y float64) (float64, error) { return 1 / x, nil },
			1, 0, 2, 1e-5,
			func(x float64) float64 { return math.Log(x) },
			nil},
		{func(x, y float64) (float64, error) { return y, nil },
			0, 1, 2, 1e-5,
			func(x float64) float64 { return math.Pow(math.E, x) },
			nil},
	}

	for i := range testEuler {
		test := &testEuler[i]
		actual, err := Euler(test.f, test.x_i, test.y_i, test.to, test.eps)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}

		act, err := actual(test.x_i)

		if err != nil {
			t.Log("error !!!!\n")
			t.FailNow()
		}

		if !EpsEqual(act, test.expected(test.x_i), 1e-3) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, act, test.expected(test.x_i))
		}

		act, err = actual(test.to - test.eps*5)

		if err != nil {
			t.Logf("euler: %v\n", err)
			t.FailNow()
		}

		tmp := test.expected(test.to - test.eps)

		if !EpsEqual(act, tmp, 1e-3) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, act, tmp)
		}

	}
}
