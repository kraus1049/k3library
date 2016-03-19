package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type inverseTest struct {
	mat      [][]float64
	expected [][]float64
	err      error
}

func TestInverse(t *testing.T) {
	var testInverse = []inverseTest{
		{[][]float64{{1, 1}, {3, 4}},
			[][]float64{{4, -1}, {-3, 1}},
			nil},
	}

	for i := range testInverse {
		test := &testInverse[i]

		actual, err := Inverse(test.mat)

		if MatEpsEqual(actual, test.expected, 1e-8) {
			t.Errorf("%v: actual = %v,expected = %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}

}
