package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type DetTest struct {
	mat      Mat
	expected float64
}

func TestDet(t *testing.T) {
	var testDet = []DetTest{
		{NewMatSet([][]float64{{1, 2}, {3, 4}}),
			-2,
		},
		{NewMatSet([][]float64{{9, 2, 5}, {6, 2, 4}, {10, 123, 2}}),
			-746,
		},
		{NewMatSet([][]float64{{12.1, 2.3, 3.4, 4.5}, {5.6, 6.7, 7.8, 8.9}, {1.3, 2.4, 3.5, 4.6}, {5.7, 68, 7.9, 8.9}}),
			3150.1326999993,
		},
	}

	for i := range testDet {
		test := &testDet[i]

		_, u, _, sgn, _ := LUDecomp(test.mat)
		actual := Det(u, sgn)

		if !EpsEqual(actual, test.expected, 1e-8) {
			t.Errorf("actual = %v, expected = %v\n", actual, test.expected)
		}
	}

}
