package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"reflect"
	"testing"
)

type vTest struct {
	a, b     []float64
	expected []float64
	err      error
}

type mProTest struct {
	a, b     [][]float64
	expected [][]float64
	err      error
}

type sumTest struct {
	xs       []interface{}
	expected interface{}
	err      error
}

func TestSum(t *testing.T) {
	v1 := NewVecSet([]float64{1, 2})
	v2 := NewVecSet([]float64{2, 4})
	v3 := NewVecSet([]float64{3, 6})
	v4 := NewVecSet([]float64{1, 2, 3})

	m1 := NewMatSet([][]float64{{1, 2}, {3, 4}})
	m2 := NewMatSet([][]float64{{2, 4}, {6, 8}})
	m3 := NewMatSet([][]float64{{3, 6}, {9, 12}})
	m4 := NewMatSet([][]float64{{1, 2, 3}, {4, 5, 6}})

	var testSum = []sumTest{
		{[]interface{}{1.0},
			1.0,
			nil},
		{[]interface{}{1.0, 2.0},
			3.0,
			nil},
		{[]interface{}{1.0, 2.0, 3.0},
			6.0,
			nil},
		{[]interface{}{v1, v1},
			v2,
			nil},
		{[]interface{}{v1, v1, v1},
			v3,
			nil},
		{[]interface{}{m1, m1},
			m2,
			nil},
		{[]interface{}{m1, m1, m1},
			m3,
			nil},
		{[]interface{}{1.0, v1},
			-1.0,
			ErrInvalid},
		{[]interface{}{1.0, v1},
			-1.0,
			ErrInvalid},
		{[]interface{}{v1, m1},
			-1.0,
			ErrInvalid},
		{[]interface{}{v1, v4},
			-1.0,
			ErrInvalid},
		{[]interface{}{m1, m4},
			-1.0,
			ErrInvalid},
	}

	for i := range testSum {
		test := &testSum[i]
		actual, err := Sum(test.xs...)

		switch actual.(type) {
		case float64:
			if actual != test.expected {
				t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
			} else if err != test.err {
				t.Errorf("%v : actual = %v, expected = %v\n", i, err, test.err)
			}
		case Vec, Mat:
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
			} else if err != test.err {
				t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
			}
		case nil:
			if err != test.err {
				t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
			}
		}

	}
}

func TestVSub(t *testing.T) {
	var testVSub = []vTest{
		{[]float64{1, 2},
			[]float64{1, 1},
			[]float64{0, 1},
			nil,
		},
	}

	for i := range testVSub {
		test := &testVSub[i]
		actual, err := VSub(test.a, test.b)

		if !SliceEpsEqual(actual, test.expected, 1e-8) {
			t.Errorf("%v: actual = %v expected = %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v: actual = %v expected = %v\n", i, err, test.err)
		}
	}
}

func TestMPro(t *testing.T) {
	var testMPro = []mProTest{
		{[][]float64{{1, 1}, {1, 1}},
			[][]float64{{1, 1}, {1, 1}},
			[][]float64{{2, 2}, {2, 2}},
			nil,
		},
		{[][]float64{{1, 2}, {3, 4}},
			[][]float64{{5, 6}, {7, 8}},
			[][]float64{{19, 22}, {43, 50}},
			nil,
		},
		{[][]float64{{1, 1, 1}, {1, 1, 1}},
			[][]float64{{1, 1}, {1, 1}, {1, 1}},
			[][]float64{{3, 3}, {3, 3}},
			nil,
		},
	}

	for i := range testMPro {
		test := &testMPro[i]
		actual, err := MPro(test.a, test.b)

		if !MatEpsEqual(actual, test.expected, 1e-8) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		} else if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}

func TestMProInvalid(t *testing.T) {
	var testMProInvalid = []mProTest{
		{[][]float64{{1, 1}},
			[][]float64{{1, 1}},
			[][]float64{{1, 1}, {1, 1}},
			ErrInvalid,
		},
		{[][]float64{{1, 1}},
			[][]float64{{1, 1}},
			[][]float64{{1, 1}, {1, 1}},
			ErrInvalid,
		},
	}

	for i := range testMProInvalid {
		test := &testMProInvalid[i]
		_, err := MPro(test.a, test.b)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}
