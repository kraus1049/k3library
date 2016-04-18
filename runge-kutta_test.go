package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"math"
	"runtime"
	"testing"
)

type rungeKuttaTest struct {
	f        FNCVec
	x_i      float64
	y_i      Vec
	to, eps  float64
	expected func(float64) float64
	err      error
}

func TestRungeKutta(t *testing.T) {
	fv1 := NewFNCVecSet(func(x float64, y Vec) (float64, error) { return 2 * x, nil })
	fv2 := NewFNCVecSet(func(x float64, y Vec) (float64, error) { return 1 / x, nil })
	fv3 := NewFNCVecSet(func(x float64, y Vec) (float64, error) { return y.At(0), nil })
	fv4 := NewFNCVecSet(func(x float64, y Vec) (float64, error) { return y.At(1), nil },
		func(x float64, y Vec) (float64, error) { return y.At(1) + 2*y.At(0), nil })
	fv5 := NewFNCVecSet(func(x float64, y Vec) (float64, error) { return y.At(1), nil },
		func(x float64, y Vec) (float64, error) { return -2*y.At(1) - 5*y.At(0), nil })

	var testRungeKutta = []rungeKuttaTest{
		{fv1,
			0, NewVecSet(0), 1, 1e-4,
			func(x float64) float64 { return math.Pow(x, 2) },
			nil},
		{fv2,
			1, NewVecSet(0), 2, 1e-4,
			func(x float64) float64 { return math.Log(x) },
			nil},
		{fv3,
			0, NewVecSet(1), 2, 1e-5,
			func(x float64) float64 { return math.Pow(math.E, x) },
			nil},
		{fv4,
			0, NewVecSet(2, -1), 0.01, 1e-5,
			func(x float64) float64 { return math.Pow(math.E, x) + math.Pow(math.E, -2*x) },
			nil},
		{fv5,
			0, NewVecSet(1, 1), 1, 1e-5,
			func(x float64) float64 { return math.Pow(math.E, -x) * (math.Cos(2*x) + math.Sin(2*x)) },
			nil},
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan bool, len(testRungeKutta))

	for i := range testRungeKutta {
		go func(i int) {
			defer func() {
				ch <- true
			}()

			test := &testRungeKutta[i]
			actual, err := RungeKutta(test.f, test.x_i, test.y_i, test.to, test.eps)

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
		}(i)

	}

	for i := 0; i < len(testRungeKutta); i++ {
		<-ch
	}
}
