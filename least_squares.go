package k3library

import (
	"math"
)

func LeastSquares(data Mat, degree int) (Vec, error) {
	if degree < 0 {
		return NewVec(0), ErrInvalid
	}

	a := NewMat(data.col, degree+1)
	b := NewVec(data.col)

	for i := 0; i < a.col; i++ {
		for j := 0; j < a.row; j++ {
			var tmp float64
			if j == 0 {
				tmp = 1.0
			} else {
				tmp = math.Pow(data.At(i, 0), float64(j))
			}

			a.Write(i, j, tmp)
		}

		b.Write(i, data.At(i, data.row-1))
	}

	ta := a.Transpose()
	tmp, _ := Pro(ta, a)
	if tmp2, ok := tmp.(Mat); ok {
		tmp3, err := tmp2.Inverse()
		if err != nil {
			return b, err
		}

		tmp4, _ := Pro(tmp3, ta, b)

		if ans, ok := tmp4.(Vec); ok {
			return ans, nil
		} else {
			return b, ErrInvalid
		}

	} else {
		return b, ErrInvalid
	}

}
