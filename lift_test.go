package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"testing"
)

func TestLift(t *testing.T) {
	num := 1.0
	f := func(x float64) float64 { return x }
	g := Lift(f)

	actual, err := g(num)

	if actual != num {
		t.Errorf("actual = %v, expected = %v\n", actual, num)
	} else if err != nil {
		t.Errorf("actual = %v, expected = %v\n", err, nil)
	}
}
