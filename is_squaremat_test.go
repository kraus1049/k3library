package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type is_squarematTest struct {
	mat      Mat
	expected bool
}

func TestIsSquareMat(t *testing.T) {
	var test_is_squaremat = []is_squarematTest{
		{NewMatSet([][]float64{{1, 1}, {1, 1}}),
			true},
		{NewMatSet([][]float64{{1, 1}, {1, 1}, {1, 1}}),
			false},
	}

	for i := range test_is_squaremat {
		test := &test_is_squaremat[i]
		actual := IsSquareMat(test.mat)

		if actual != test.expected {
			t.Errorf("%v: actual = %v, expected = %v", i, actual, test.expected)
		}
	}
}
