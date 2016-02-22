package k3library

import "testing"

type epsequalTest struct {
	a, b, eps float64
	expected  bool
}

func TestEpsequal(t *testing.T) {

	// a,b,eps expected
	testepsequall := []epsequalTest{
		{1, 2, 3, true},
		{1, 2, 0, false},
		{1e-1, 11e-2, 1e-1, true},
		{1e-1, 11e-2, 1e-2, true},
		{1e-1, 11e-2, 1e-3, false},
	}

	for i := range testepsequall {
		test := &testepsequall[i]
		actual := Epsequal(test.a, test.b, test.eps)

		if actual != test.expected {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		}
	}
}
