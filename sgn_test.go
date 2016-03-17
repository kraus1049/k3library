package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type sgnTest struct {
	x, expected float64
	err         error
}

func TestSgn(t *testing.T) {
	var testsgn = []sgnTest{
		{2, 1, nil},
		{0, 0, nil},
		{-3, -1, nil},
	}

	for i := range testsgn {
		test := &testsgn[i]
		actual, err := Sgn(test.x)

		if actual != test.expected {
			t.Errorf("%v:actual = %v,expected = %v\n", actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v:actual = %v,expected = %v\n", err, test.err)
		}
	}

}
