package k3library_test

import (
	"fmt"
	. "github.com/kraus1049/k3library"
	"reflect"
	"testing"
)

type vSetTest struct {
	x        []float64
	len      int
	expected []float64
}

type mSetTest struct {
	x        [][]float64
	col, row int
	expected [][]float64
}

type fncVSetTest struct {
	fs       []func(float64, Vec) (float64, error)
	row      int
	expected []func(float64, Vec) (float64, error)
}

type newFNCVSet struct {
	fs       []func(float64, Vec) (float64, error)
	expected []func(float64, Vec) (float64, error)
}

type vWriteTest struct {
	len, idx int
	num      float64
	expected []float64
}

type mWriteTest struct {
	col, row, i, j int
	num            float64
	expected       [][]float64
}

type fncvCalc struct {
	fv       FNCVec
	x        float64
	y        Vec
	expected Vec
	err      error
}

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

func TestVSet(t *testing.T) {
	var testVSet = []vSetTest{
		{[]float64{1, 2, 3, 4, 5},
			3,
			[]float64{1, 2, 3}},
		{[]float64{1, 2, 3, 4, 5},
			5,
			[]float64{1, 2, 3, 4, 5}},
		{[]float64{1, 2, 3, 4, 5},
			7,
			[]float64{1, 2, 3, 4, 5, 0, 0}},
	}

	for i := range testVSet {
		test := &testVSet[i]
		v := NewVec(test.len)
		v.Set(test.x)

		if !reflect.DeepEqual(v.V, test.expected) {
			t.Errorf("%v: actual = %v, expected = %v\n", v.V, test.expected)
		}
	}
}

func TestNewVecSet(t *testing.T) {
	xs := []float64{0, 1, 2, 3, 4}
	v := NewVecSet(xs...)

	if !reflect.DeepEqual(v.V, xs) {
		t.Errorf("actual = %v, expected = %v\n", v.V, xs)
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

func TestNewMatSet(t *testing.T) {
	xss := [][]float64{{1, 2, 3}, {4, 5, 6}}
	m := NewMatSet(xss)

	flag := false
	for i := 0; i < m.Col; i++ {
		if !reflect.DeepEqual(m.M[i].V, xss[i]) {
			flag = true
			break
		}
	}

	if flag {
		t.Errorf("actual = %v, expected = %v\n", m, xss)
	}
}

func TestMSet(t *testing.T) {
	var testMSet = []mSetTest{
		{[][]float64{{1, 2}, {3, 4}, {5, 6}},
			3, 2,
			[][]float64{{1, 2}, {3, 4}, {5, 6}}},
		{[][]float64{{1, 2}, {3, 4}, {5, 6}},
			2, 2,
			[][]float64{{1, 2}, {3, 4}}},
		{[][]float64{{1, 2}, {3, 4}, {5, 6}},
			4, 2,
			[][]float64{{1, 2}, {3, 4}, {5, 6}, {0, 0}}},
		{[][]float64{{1, 2}, {3, 4}, {5, 6}},
			3, 1,
			[][]float64{{1}, {3}, {5}}},
		{[][]float64{{1, 2}, {3, 4}, {5, 6}},
			3, 3,
			[][]float64{{1, 2, 0}, {3, 4, 0}, {5, 6, 0}}},
	}

	for i := range testMSet {
		test := &testMSet[i]
		m := NewMat(test.col, test.row)
		m.Set(test.x)

		flag := false
		for j := 0; j < m.Col; j++ {
			if !reflect.DeepEqual(m.M[j].V, test.expected[j]) {
				flag = true
				break
			}

		}

		if flag {
			t.Errorf("%v: actual = %v, expected = %v\n", i, m.M, test.expected)
		}
	}
}

func TestNewFNCVec(t *testing.T) {
	for i := 0; i < 10; i++ {
		actual := NewFNCVec(i)

		if len(actual.F) != i {
			t.Errorf("%v: actual = %v, expected = %v\n", i, len(actual.F), i)
		} else if actual.Row != i {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual.Row, i)
		}
	}
}

func TestFNCVSet(t *testing.T) {
	// y := NewVecSet([]float64{1, 2})
	fv1 := []func(float64, Vec) (float64, error){
		func(x float64, y Vec) (float64, error) { return x * y.At(0), nil },
	}

	fv2 := []func(float64, Vec) (float64, error){
		func(x float64, y Vec) (float64, error) { return x * y.At(0), nil },
		func(x float64, y Vec) (float64, error) { return x + y.At(0) + y.At(1), nil },
	}

	var testFNCVSet = []fncVSetTest{
		{fv1, 1, fv1},
		{fv2, 2, fv2},
	}

	for i := range testFNCVSet {
		test := &testFNCVSet[i]
		fv := NewFNCVec(test.row)
		fv.Set(test.fs...)

		if !reflect.DeepEqual(fv.F, test.fs) {
			t.Errorf("%v: actual = %v, expected = %v", i, fv.F, test.fs)
		}
	}
}

func TestNewFNCVSet(t *testing.T) {
	// y := NewVecSet([]float64{1, 2})
	fv1 := []func(float64, Vec) (float64, error){
		func(x float64, y Vec) (float64, error) { return x * y.At(0), nil },
	}

	fv2 := []func(float64, Vec) (float64, error){
		func(x float64, y Vec) (float64, error) { return x * y.At(0), nil },
		func(x float64, y Vec) (float64, error) { return x + y.At(0) + y.At(1), nil },
	}

	var testNewFNCVSet = []newFNCVSet{
		{fv1, fv1},
		{fv2, fv2},
	}

	for i := range testNewFNCVSet {
		test := &testNewFNCVSet[i]
		fv := NewFNCVecSet(test.fs...)

		if !reflect.DeepEqual(fv.F, test.fs) {
			t.Errorf("%v: actual = %v, expected = %v", i, fv.F, test.fs)
		}
	}

}

func TestVAt(t *testing.T) {
	v := NewVec(5)
	expected := []float64{0, 1, 2, 3, 4}

	v.Set(expected)

	actual := make([]float64, 0)

	for i := range v.V {
		actual = append(actual, v.At(i))
	}

	if !reflect.DeepEqual(actual, v.V) {
		t.Errorf("actual = %v, expected = %v\n", actual, expected)
	}
}

func TestMAt(t *testing.T) {
	m := NewMat(3, 2)
	expected := [][]float64{{1, 2}, {3, 4}, {5, 6}}

	m.Set(expected)

	actual := make([][]float64, 3)

	for i := range m.M {
		for j := 0; j < m.Row; j++ {
			actual[i] = append(actual[i], m.At(i, j))
		}
	}

	flag := false
	for i := range m.M {
		if !reflect.DeepEqual(m.M[i].V, actual[i]) {
			flag = true
			break
		}
	}

	if flag {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

func TestFNCVAt(t *testing.T) {
	y := NewVecSet(1, 2, 3, 4)
	f1 := func(x float64, y Vec) (float64, error) { return x + y.At(0) + y.At(3), nil }
	f2 := func(x float64, y Vec) (float64, error) { return x * y.At(1), nil }
	f3 := func(x float64, y Vec) (float64, error) { return x - y.At(2), nil }
	f4 := func(x float64, y Vec) (float64, error) { return x + 2*y.At(3), nil }

	fvs := [][]func(float64, Vec) (float64, error){{f1, f2}, {f3, f4}}

	for i := range fvs {
		fv := NewFNCVecSet(fvs[i]...)

		for j := 0; j < fv.Row; j++ {
			tmp := fv.At(j)
			actual, _ := tmp(1, y)
			expected, _ := fvs[i][j](1, y)

			if actual != expected {
				t.Errorf("(%v,%v): actual = %v, expected = %v\n", i, j, actual, expected)
			}
		}
	}
}

func TestVWrite(t *testing.T) {
	var testVWrite = []vWriteTest{
		{5, 0, 1, []float64{1, 0, 0, 0, 0}},
		{5, 3, 2, []float64{0, 0, 0, 2, 0}},
		{5, -1, 3, []float64{3, 3, 3, 3, 3}},
	}

	for i := range testVWrite {
		test := &testVWrite[i]
		v := NewVec(test.len)
		v.Write(test.idx, test.num)

		if !reflect.DeepEqual(test.expected, v.V) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, v.V, test.expected)
		}

	}
}

func TestMWrite(t *testing.T) {
	var testMWrite = []mWriteTest{
		{2, 3, 0, 0, 1, [][]float64{{1, 0, 0}, {0, 0, 0}}},
		{2, 3, -1, 0, 1, [][]float64{{1, 0, 0}, {1, 0, 0}}},
		{2, 3, 0, -1, 1, [][]float64{{1, 1, 1}, {0, 0, 0}}},
		{2, 3, -1, -1, 1, [][]float64{{1, 0, 0}, {0, 1, 0}}},
		{3, 3, -1, -1, 1, [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
	}

	for i := range testMWrite {
		test := &testMWrite[i]
		m := NewMat(test.col, test.row)
		m.Write(test.i, test.j, test.num)

		flag := false
		for j := 0; j < m.Col; j++ {
			if !reflect.DeepEqual(m.M[j].V, test.expected[j]) {
				flag = true
				break
			}
		}

		if flag {
			t.Errorf("%v: actual = %v, expected = %v\n", i, m.M, test.expected)
		}

	}

}

func TestVCopy(t *testing.T) {
	v1 := Vec{[]float64{1, 2, 3}, 3}
	v2 := v1.Copy()

	if !reflect.DeepEqual(v1, v2) {
		t.Errorf("actual = %v, expected = %v\n", v1, v2)
	}

	v2.Write(0, 100)

	if reflect.DeepEqual(v1, v2) {
		t.Errorf("v1(%v) should not be equal v2(%v)\n", v1, v2)
	}

}

func TestMCopy(t *testing.T) {
	m1 := NewMatSet([][]float64{{1, 2}, {3, 4}})
	m2 := m1.Copy()

	if !reflect.DeepEqual(m1, m2) {
		t.Errorf("actual = %v, expected = %v\n", m1, m2)
	}

	m2.Write(0, 0, 100)

	if reflect.DeepEqual(m1, m2) {
		t.Errorf("m1(%v) should be not equal m2(%v)\n", m1, m2)
	}

}

func TestFNCVCalc(t *testing.T) {
	f1 := func(x float64, y Vec) (float64, error) { return x * y.At(0), nil }
	f2 := func(x float64, y Vec) (float64, error) { return x * y.At(1), nil }

	var testFNCVCalc = []fncvCalc{
		{NewFNCVecSet(f1, f2), 1,
			NewVecSet(1, 2),
			NewVecSet(1, 2), nil},
	}

	for i := range testFNCVCalc {
		test := &testFNCVCalc[i]

		actual, err := test.fv.Calc(test.x, test.y)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		} else if !VecEpsEqual(actual, test.expected, 1e-5) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		}
	}
}

func ExamplePrintVec() {
	v := NewVec(3)
	fmt.Print(v)
	// Output:
	//
	// [0 0 0]
}

func ExampleMat1() {
	m := NewMat(3, 3)
	fmt.Print(m)
	// Output:
	//
	// | 0 0 0 |
	// | 0 0 0 |
	// | 0 0 0 |
}
