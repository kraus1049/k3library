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

		actual, err := Inverse(test.mat)

		if pro, _ := Pro(actual, test.mat); !isIdentityMat(pro.(Mat), 1e-8) {
			t.Errorf("%v: want identityMat, but actual = %v\n", i, pro)
		} else if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}

}

func BenchmarkInverse(b *testing.B) {
	mat := NewMat(5, 5)
	for i := 0; i < mat.Col; i++ {
		for j := 0; j < mat.Row; j++ {
			mat.Write(i, j, float64(i+j))
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Inverse(mat)
	}
}

func isIdentityMat(mat Mat, eps float64) bool {
	if !IsSquareMat(mat) {
		return false
	}

	for i := 0; i < mat.Col; i++ {
		if !EpsEqual(mat.At(i, i), 1, eps) {
			return false
		}
	}

	return true
}
