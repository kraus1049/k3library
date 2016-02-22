package k3library

import(
	"testing"
)

type csv2funcTest struct{
	filepath string
	xy [][2]float64
	ok bool
}

func TestCsv2func(t *testing.T){
	var testcsv2func  = []csv2funcTest{
		{"./testdata/test.CSV",
		{},
	true},
	}

	testfile := "./testdata/test.CSV"

	f,ok := Csv2func(testfile)


}
