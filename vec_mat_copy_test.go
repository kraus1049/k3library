package k3library

import (
	"reflect"
	"testing"
)

type vec_copyTest struct {
	vec      []float64
	expected []float64
}

type mat_copyTest struct {
	mat      [][]float64
	expected [][]float64
}

func TestVecCopy(t *testing.T) {
	var testveccopy = []vec_copyTest{
		{[]float64{1, 2, 3, 4, 5},
			[]float64{1, 2, 3, 4, 5}},
	}

	for i := range testveccopy {
		test := &testveccopy[i]
		actual := Vec_copy(test.vec)

		test.vec[0] = 100

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%v:actual = %v,expected = %v\n", i, actual, test.expected)
		}else if reflect.DeepEqual(actual,test.vec){
			t.Errorf("%v:I want actual = %v is not equal expected = %v\n", i, actual, test.vec)
		}
	}
}

func TestMatCopy(t *testing.T) {
	var testmatcopy = []mat_copyTest{
		{[][]float64{{1, 2}, {3, 4}},
			[][]float64{{1, 2}, {3, 4}}},
	}

	for i := range testmatcopy {
		test := &testmatcopy[i]
		actual := Mat_copy(test.mat)

		test.mat[0][0] = 100

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%v:actual = %v,expected = %v\n", i, actual, test.expected)
		}else if reflect.DeepEqual(actual,test.mat){
			t.Errorf("%v:I want actual = %v is not equal expected = %v\n",i,actual,test.mat)
		}
	}
}
