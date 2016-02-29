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
}

func TestCsv2func(t *testing.T) {
	var testcsv2func = []csv2funcTest{
		{"./testdata/F0000CH1.CSV",
			goscilloscope.GOscilloscope,
			[][2]float64{{-0.0000543, -0.04}},
			error{nil},
		},
	}

	for i := range testcsv2func {
		test := &testcsv2func[i]

		f, func_err := Csv2func(test.filepath, test.parser)

		for j := range test.expected_xy {
			x, y := expected_xy[j]

			actual_y, err := f(x)

			if y != actual_y {
				t.Errorf("%v,%v: actual (%v,%v), expected (%v,%v)", i, j, x, actual_y, x, y)
			} else if err != test.errs[j] {
				t.Errorf("%v,%v: actual %v, expected %v", i, j, err, test.errs[j])
			}
		}
	}
}
