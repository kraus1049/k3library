package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type inverseTest struct {
	mat Mat
	err error
}

func TestInverse(t *testing.T) {
	var testInverse = []inverseTest{
		{NewMatSet([][]float64{{1, 1}, {3, 4}}),
			nil},
		{NewMatSet([][]float64{{45, 67, 89}, {111, 121, 23}, {34, 4, 831}}),
			nil},
	}

	for i := range testInverse {
		test := &testInverse[i]

		actual, err := test.mat.Inverse()

		if pro, _ := Pro(actual, test.mat); !isIdentityMat(pro.(Mat),
			1e-8) {
			t.Errorf("%v: want identityMat, but actual = %v\n", i,
				pro)
		} else if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}

}

func BenchmarkInverse(b *testing.B) {
	m := [][]float64{
		{1, 3, 5, 7, 9},
		{2, 3, 5, 6, 7},
		{7, 6, 43, 6, 2},
		{1, 2, 7, 9, 4},
		{90, 56, 42, 5, 6},
	}

	mat := NewMatSet(m)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := mat.Inverse()

		if err != nil {
			b.Logf("%v", err)
		}
	}
}

func isIdentityMat(mat Mat, eps float64) bool {
	if !mat.IsSquareMat() {
		return false
	}

	for i := 0; i < mat.Col(); i++ {
		if !EpsEqual(mat.At(i, i), 1, eps) {
			return false
		}
	}

	return true
}
