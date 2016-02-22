package k3library

import (
	"math"
	"testing"
)

type zbrakTest struct {
	f          func(x float64) float64
	start, end float64
	num        float64
	div_n      int
	expected   []float64
	ok         bool
}

func TestZbrak(t *testing.T) {
	// f,start, end , num, div_n,expected ok
	testzbrak := []zbrakTest{
		{func(x float64) float64 { return x },
			-5, 5, 0, 10,
			[]float64{0},
			true},
		{func(x float64) float64 { return math.Sin(x) },
			-2*math.Pi - 1, 2*math.Pi + 1, 0, 100000,
			[]float64{-2 * math.Pi, -math.Pi, 0, math.Pi, 2 * math.Pi},
			true},
	}

	for i := range testzbrak {
		test := &testzbrak[i]

		actual, ok := Zbrak(test.start, test.end, test.num, test.div_n, test.f)

		if ok != test.ok {
			t.Errorf("actual = %v,expected = %v\n", ok, test.ok)
		}

		if len(actual) != len(test.expected) {
			t.Errorf("len(actual) = %v,len(expected) = %v", len(actual), len(test.expected))
		}

		for _, item := range test.expected {
			flag := false
			for _, test_item := range actual {
				if test_item[0] <= item && item <= test_item[1] {
					flag = true
				}
			}
			if flag == false {
				t.Errorf("%v isn't included in actual\n", test.expected)
			}
		}
	}
}

func TestZbrakInvalidArgument(t *testing.T) {
	// f,start, end , num, div_n,expected ok
	testzbrakinvalidargument := []zbrakTest{
		{func(x float64) float64 { return x },
			1, 1, 0, 100, []float64{},
			false},
		{func(x float64) float64 { return x },
			1, 2, 0, -10, []float64{},
			false},
	}

	for i := range testzbrakinvalidargument {
		test := &testzbrakinvalidargument[i]
		_, ok := Zbrak(test.start, test.end, test.num, test.div_n, test.f)

		if ok != test.ok {
			t.Errorf("actual = %v,expected = %v\n", ok, test.ok)
		}
	}
}
