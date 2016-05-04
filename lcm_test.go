package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type lcmTest struct {
	x, y, expected int
}

func TestLCM(t *testing.T) {
	var testLCM = []lcmTest{
		{3, 4, 12},
		{5, 4, 20},
		{4, 6, 12},
		{32, 55, 1760},
	}

	for i := range testLCM {
		test := &testLCM[i]
		actual := LCM(test.x, test.y)

		if actual != test.expected {
			t.Errorf("%v: actual = %v, expected = %v", i, actual, test.expected)
		}
	}
}
