package k3library

func VSub(x, y []float64) ([]float64, error) {
	if len(x) != len(y) {
		return nil, ErrInvalid
	}

	ans := make([]float64, len(x))

	for i := range x {
		ans[i] = x[i] - y[i]
	}

	return ans, nil
}

func MPro(x, y [][]float64) ([][]float64, error) {

	if len(x[0]) != len(y) {
		return nil, ErrInvalid
	}

	ans := makeMat(len(x), len(y[0]))

	for i := range x {
		for j := range y[0] {
			for k := range x[0] {
				ans[i][j] += x[i][k] * y[k][j]
			}
		}
	}

	return ans, nil
}

func Sum(xs ...interface{}) (interface{}, error) {
	if !allTypeEqual(xs...) {
		return nil, ErrInvalid
	}

	if len(xs) == 1 {
		if x, ok := xs[0].(float64); ok {
			return x, nil
		} else if v, ok := xs[0].(Vec); ok {
			return v, nil
		} else if m, ok := xs[0].(Mat); ok {
			return m, nil
		}

	} else if len(xs) >= 2 {
		switch xs[0].(type) {
		case float64:
			x, _ := xs[0].(float64)
			y, _ := xs[1].(float64)

			xs = append(xs[2:], x+y)
		case Vec:
			x, _ := xs[0].(Vec)
			y, _ := xs[1].(Vec)
			if x.Row != y.Row {
				return nil, ErrInvalid
			}

			plus := NewVec(x.Row)

			for i := 0; i < x.Row; i++ {
				tmp := x.At(i) + y.At(i)
				plus.Write(i, tmp)
			}

			xs = append(xs[2:], plus)

		case Mat:
			x, _ := xs[0].(Mat)
			y, _ := xs[1].(Mat)

			if x.Col != y.Col || x.Row != y.Row {
				return nil, ErrInvalid
			}

			plus := NewMat(x.Col, x.Row)

			for i := 0; i < x.Col; i++ {
				for j := 0; j < x.Row; j++ {
					tmp := x.At(i, j) + y.At(i, j)
					plus.Write(i, j, tmp)
				}
			}

			xs = append(xs[2:], plus)
		}

		ans, err := Sum(xs...)
		return ans, err

	}
	return nil, ErrInvalid
}
