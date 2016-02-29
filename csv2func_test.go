package k3library

import (
	"github.com/kraus1049/goscilloscope"
	"testing"
)

type csv2funcTest struct {
	filepath    string
	parser      func(string) (map[string]string, [][2]float64, error)
	expected_xy [][2]float64
	errs        []error
	func_err    error
}

func TestCsv2func(t *testing.T) {
	var testcsv2func = []csv2funcTest{
		{"./testdata/F0000CH1.CSV",
			goscilloscope.GOscilloscope,
			[][2]float64{
				{-0.000055, -0.04},
				{-0.0000543, -0.04},
				{-0.0000454, -0.08},
				{0.0000145, 4.24},
				{0.00002483, 5.0},
				{0.0001924, -0.04},
				{0.0001949, -0.04},
			},
			[]error{nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
			},
			nil,
		},
	}

	for i := range testcsv2func {
		test := &testcsv2func[i]

		f, func_err := Csv2func(test.filepath, test.parser)

		if func_err != test.func_err {
			t.Errorf("func_err,%v: actual %v,expected %v", i, func_err, test.func_err)
		}

		for j := range test.expected_xy {
			tmp := test.expected_xy[j]
			x, y := tmp[0], tmp[1]

			actual_y, err := f(x)

			if y != actual_y {
				t.Errorf("%v,%v: actual (%v,%v), expected (%v,%v)", i, j, x, actual_y, x, y)
			} else if err != test.errs[j] {
				t.Errorf("%v,%v: actual %v, expected %v", i, j, err, test.errs[j])
			}
		}
	}
}

func BenchmarkCsv2func(b *testing.B) {
	f, _ := Csv2func("./testdata/F0000CH1.CSV", goscilloscope.GOscilloscope)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(-0.000055)
		f(-0.000044)
		f(-0.000033)
		f(-0.000022)
		f(-0.000011)
		f(0)
		f(0.000011)
		f(0.000022)
		f(0.000033)
		f(0.000044)
		f(0.000055)
		f(0.000066)
		f(0.000077)
		f(0.000088)
		f(0.000099)
		f(0.000111)
	}
}
