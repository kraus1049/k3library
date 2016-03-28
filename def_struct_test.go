package k3library_test

import (
	"fmt"
	. "github.com/kraus1049/k3library"
	"testing"
)

func TestNewVec(t *testing.T) {
	for i := 0; i < 10; i++ {
		v := NewVec(i)
		if len(v.V) != i {
			t.Errorf("%v: actual = %v, expected = %v\n", i, i, len(v.V))
		} else if v.Row != i {
			t.Errorf("%v: actual = %v, expected = %v\n", i, i, v.Row)
		}
	}
}

func TestNewMat(t *testing.T) {
	for i := 1; i < 10; i++ {
		for j := 0; j < 10; j++ {
			m := NewMat(i, j)
			if len(m.M) != i {
				t.Errorf("%v,%v: actual = %v, expected = %v\n", i, j, len(m.M), i)
			} else if len(m.M[0].V) != j {
				t.Errorf("%v,%v: actual = %v, expected = %v\n", i, j, len(m.M[0].V), j)
			} else if m.Col != i {
				t.Errorf("%v,%v: actual = %v, expected = %v\n", i, j, m.Col, i)
			} else if m.Row != j {
				t.Errorf("%v,%v: actual = %v, expected = %v\n", i, j, m.Row, j)
			}
		}
	}
}

func ExamplePrintVec() {
	v := NewVec(3)
	fmt.Print(v)
	// Output:
	// [0 0 0]
}

func ExampleMat1() {
	m := NewMat(3, 3)
	fmt.Print(m)
	// Output:
	// | 0 0 0 |
	// | 0 0 0 |
	// | 0 0 0 |

}
