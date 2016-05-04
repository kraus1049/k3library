package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

type gcdTest struct {
	x, y, expected int
}

func TestGCD(t *testing.T) {
	var testGCD = []gcdTest{
		{12, 18, 6},
		{12420, 18930, 30},
		{12423, 18931, 1},
	}

	for i := range testGCD {
		test := &testGCD[i]

		actual := GCD(test.x, test.y)

		if actual != test.expected {
			t.Errorf("%v: actual = %v, expected = %v", i, actual, test.expected)
		}
	}
}
