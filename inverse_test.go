package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type inverseTest struct {
	mat [][]float64
	err error
}

func TestInverse(t *testing.T) {
	var testInverse = []inverseTest{
		{[][]float64{{1, 1}, {3, 4}},
			nil},
		{[][]float64{{45, 67, 89}, {111, 121, 23}, {34, 4, 831}},
			nil},
	}

	for i := range testInverse {
		test := &testInverse[i]

		actual, err := Inverse(test.mat)

		if pro, _ := MPro(actual, test.mat); !isIdentityMat(pro, 1e-8) {
			t.Errorf("%v: want identityMat, but actual = %v\n", i, pro)
		} else if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}

}

func isIdentityMat(mat [][]float64, eps float64) bool {
	if !IsSquareMat(mat) {
		return false
	}

	for i := range mat {
		if !EpsEqual(mat[i][i], 1, eps) {
			return false
		}
	}

	return true
}
