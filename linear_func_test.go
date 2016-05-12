package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type linearFuncTest struct {
	vec Vec
	x_y [][2]float64
}

func TestLinearFunc(t *testing.T) {
	testLinearFunc := []linearFuncTest{
		{NewVecSet(1),
			[][2]float64{{0, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1}}},
		{NewVecSet(1, 2),
			[][2]float64{{1, 3}, {2, 5}, {3, 7}, {4, 9}, {5, 11}}},
		{NewVecSet(1, 2, 3),
			[][2]float64{{1, 6}, {2, 17}, {3, 34}, {4, 57}, {5, 86}}},
		{NewVecSet(10, 9, 8, 7, 6, 5, 4, 3, 2, 1),
			[][2]float64{{1, 55},
				{2, 2036},
				{3, 44281},
				{4, 466030},
				{5, 3051755}}},
	}

	for i := range testLinearFunc {
		test := &testLinearFunc[i]

		f := LinearFunc(test.vec)

		for j := range test.x_y {
			y := f(test.x_y[j][0])

			if yy := test.x_y[j][1]; !EpsEqual(y, yy, 1e-5) {
				t.Errorf("%v: actual = %v, expected = %v\n", i, y, yy)
			}
		}
	}
}
