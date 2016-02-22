package k3library

import (
	"testing"
)

type sgnTest struct {
	x, expected float64
	ok          bool
}

func TestSgn(t *testing.T) {
	var testsgn = []sgnTest{
		{2, 1, true},
		{0, 0, true},
		{-3, -1, true},
	}

	for i := range testsgn {
		test := &testsgn[i]
		actual, ok := Sgn(test.x)

		if actual != test.expected {
			t.Errorf("%v:actual = %v,expected = %v\n", actual, test.expected)
		} else if ok != test.ok {
			t.Errorf("%v:actual = %v,expected = %v\n", ok, test.ok)
		}
	}

}
