package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"reflect"
	"testing"
)

type transposeTest struct {
	mat, expected Mat
}

func TestTranspose(t *testing.T) {
	var testTranspose = []transposeTest{
		{NewMatSet([][]float64{{1, 2}, {3, 4}}),
			NewMatSet([][]float64{{1, 3}, {2, 4}})},
		{NewMatSet([][]float64{{1, 2, 3}, {4, 5, 6}}),
			NewMatSet([][]float64{{1, 4}, {2, 5}, {3, 6}})},
		{NewMatSet([][]float64{{1, 2}, {3, 4}, {5, 6}, {7, 8}}),
			NewMatSet([][]float64{{1, 3, 5, 7}, {2, 4, 6, 8}})},
	}

	for i := range testTranspose {
		test := &testTranspose[i]
		actual := test.mat.Transpose()

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		} else if !reflect.DeepEqual(actual.Transpose(), test.mat) {
			t.Errorf("%v: Transpose(Transpose(mat)) != mat", i)
		}
	}
}

func BenchmarkTranspose(b *testing.B) {
	n := 50
	m := NewMat(n, n)

	for i := 0; i < m.Col; i++ {
		for j := 0; j < m.Row; j++ {
			m.Write(i, j, float64(i+j))
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Transpose()
	}

}
