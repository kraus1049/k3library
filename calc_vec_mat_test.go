package k3library_test

import (
	. "github.com/kraus1049/k3library"
	"math"
	"reflect"
	"testing"
)

type sumTest struct {
	xs       []interface{}
	expected interface{}
	err      error
}

type sumTestInvalidArgument struct {
	xs  []interface{}
	err error
}

type proTest struct {
	xs       []interface{}
	expected interface{}
	err      error
}

type proTestInvalidArgument struct {
	xs  []interface{}
	err error
}

type subTest struct {
	x, y     interface{}
	expected interface{}
	err      error
}

type subTestInvalidArgument struct {
	x, y interface{}
	err  error
}

type vAverageTest struct {
	xs       Vec
	expected float64
}

type mAverageTest struct {
	xs       Mat
	expected float64
}

type vAbs struct {
	xs       Vec
	expected float64
}

func TestSum(t *testing.T) {
	v1 := NewVecSet(1, 2)
	v2 := NewVecSet(2, 4)
	v3 := NewVecSet(3, 6)
	v4 := NewVecSet(1, 2, 3)

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

func TestSumInvalidArgument(t *testing.T) {
	x1 := 1.0
	v1 := NewVecSet(1, 1)
	v2 := NewVecSet(1, 1, 1)
	m1 := NewMatSet([][]float64{{1, 2}, {3, 4}})
	m2 := NewMatSet([][]float64{{1, 2}, {3, 4}, {5, 6}})

	var testSumInvalidArgument = []sumTestInvalidArgument{
		{[]interface{}{x1, v1}, ErrInvalid},
		{[]interface{}{v1, x1}, ErrInvalid},
		{[]interface{}{x1, m1}, ErrInvalid},
		{[]interface{}{m1, x1}, ErrInvalid},
		{[]interface{}{v1, m1}, ErrInvalid},
		{[]interface{}{m1, v1}, ErrInvalid},
		{[]interface{}{v1, v2}, ErrInvalid},
		{[]interface{}{m1, m2}, ErrInvalid},
	}

	for i := range testSumInvalidArgument {
		test := &testSumInvalidArgument[i]
		_, err := Sum(test.xs...)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}

func TestPro(t *testing.T) {
	x1 := 2.0
	x2 := 4.0
	x3 := 8.0
	x4 := 28.0
	x5 := 17640.0
	v1 := NewVecSet(1, 2, 3)
	v2 := NewVecSet(2, 4, 6)
	v3 := NewVecSet(8, 16, 24)
	v5 := NewVecSet(14, 32, 50)
	v6 := NewVecSet(1, 2)
	m1 := NewMatSet([][]float64{{1, 2, 3}, {4, 5, 6}})
	m2 := NewMatSet([][]float64{{2, 4, 6}, {8, 10, 12}})
	m3 := NewMatSet([][]float64{{1, 2, 3}})
	m4 := NewMatSet([][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	m5 := NewMatSet([][]float64{{30, 36, 42}, {66, 81, 96}, {102, 126, 150}})
	m6 := NewMatSet([][]float64{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}})

	var testPro = []proTest{
		{[]interface{}{x1, x1}, x2, nil},
		{[]interface{}{x1, x1, x1}, x3, nil},
		{[]interface{}{x1, v1}, v2, nil},
		{[]interface{}{x1, x1, v2}, v3, nil},
		{[]interface{}{x1, m1}, m2, nil},
		{[]interface{}{v1, x1}, v2, nil},
		{[]interface{}{v1, v2}, x4, nil},
		{[]interface{}{v1, m3}, m6, nil},
		{[]interface{}{m1, x1}, m2, nil},
		{[]interface{}{m4, v1}, v5, nil},
		{[]interface{}{m4, m4}, m5, nil},
		{[]interface{}{v1, v1, m1, m4, v1, v6},
			x5, nil},
	}

	for i := range testPro {
		test := &testPro[i]
		actual, err := Pro(test.xs...)

		switch actual.(type) {
		case float64:
			if actual != test.expected {
				t.Errorf("%v: actual = %v,expected = %v\n", i, actual, test.expected)
			} else if err != test.err {
				t.Errorf("%v: actual = %v,expected = %v\n", i, actual, test.expected)
			}

		case Vec, Mat:
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("%v: actual = %v,expected = %v\n", i, actual, test.expected)
			} else if err != test.err {
				t.Errorf("%v: actual = %v,expected = %v\n", i, actual, test.expected)
			}

		case nil:
			if err != test.err {
				t.Errorf("%v: actual = %v,expected = %v\n", i, actual, test.expected)
			}
		}

	}

}

func TestProInvalidArgument(t *testing.T) {
	v1 := NewVecSet(1, 2, 3)
	v2 := NewVecSet(1, 2, 3, 4)
	m1 := NewMatSet([][]float64{{1, 2}, {3, 4}})
	m2 := NewMatSet([][]float64{{1}, {2}, {3}})

	var testProInvalidArgument = []proTestInvalidArgument{
		{[]interface{}{v1, v2}, ErrInvalid},
		{[]interface{}{v2, v1}, ErrInvalid},
		{[]interface{}{m1, m2}, ErrInvalid},
		{[]interface{}{m2, m1}, ErrInvalid},
		{[]interface{}{v1, m1}, ErrInvalid},
		{[]interface{}{m1, v1}, ErrInvalid},
	}

	for i := range testProInvalidArgument {
		test := &testProInvalidArgument[i]
		_, err := Pro(test.xs...)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}
}

func TestSub(t *testing.T) {
	x1 := 1.0
	x2 := 2.0
	v1 := NewVecSet(1, 2, 3, 4, 5, 6, 7, 8, 9)
	v2 := NewVecSet(2, 4, 6, 8, 10, 12, 14, 16, 18)
	m1 := NewMatSet([][]float64{{1, 2}, {3, 4}})
	m2 := NewMatSet([][]float64{{2, 4}, {6, 8}})
	var testSub = []subTest{
		{x2, x1, x1, nil},
		{v2, v1, v1, nil},
		{m2, m1, m1, nil},
	}

	for i := range testSub {
		test := &testSub[i]
		actual, err := Sub(test.y, test.x)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		} else if reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		}
	}
}

func TestSubInvalidArgument(t *testing.T) {
	x1 := 1.0
	v1 := NewVecSet(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m1 := NewMatSet([][]float64{{1, 2}, {3, 4}})

	var testSubInvalidArgument = []subTestInvalidArgument{
		{x1, v1, ErrInvalid},
		{v1, x1, ErrInvalid},
		{x1, m1, ErrInvalid},
		{m1, x1, ErrInvalid},
		{v1, m1, ErrInvalid},
		{m1, v1, ErrInvalid},
	}

	for i := range testSubInvalidArgument {
		test := &testSubInvalidArgument[i]
		_, err := Sub(test.x, test.y)

		if err != test.err {
			t.Errorf("%v: actual = %v, expected = %v\n", i, err, test.err)
		}
	}

}

func TestVAverage(t *testing.T) {
	var testVAverage = []vAverageTest{
		{NewVecSet(1.0), 1.0},
		{NewVecSet(1, 2), 1.5},
		{NewVecSet(1, 2, 3), 2.0},
	}

	for i := range testVAverage {
		test := &testVAverage[i]
		actual := test.xs.Average()

		if actual != test.expected {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		}

	}
}

func TestMAverage(t *testing.T) {
	var testMAverage = []mAverageTest{
		{NewMatSet([][]float64{{1}, {2}}), 1.5},
		{NewMatSet([][]float64{{1, 2}, {3, 4}}), 2.5},
		{NewMatSet([][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), 5.0},
	}

	for i := range testMAverage {
		test := &testMAverage[i]
		actual := test.xs.Average()

		if actual != test.expected {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		}
	}
}

func TestAbs(t *testing.T) {
	var testAbs = []vAbs{
		{NewVecSet(2),
			2},
		{NewVecSet(1, 2),
			math.Sqrt(5)},
		{NewVecSet(1, 2, 3),
			math.Sqrt(14)},
	}

	for i := range testAbs {
		test := &testAbs[i]
		actual := test.xs.Abs()

		if actual != test.expected {
			t.Errorf("%v: actual = %v, expected = %v\n", i, actual, test.expected)
		}
	}
}
