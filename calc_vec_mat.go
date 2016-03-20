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

	ans := make([][]float64, len(x))
	for i := range ans {
		ans[i] = make([]float64, len(y[0]))
	}

	for i := range x {
		for j := range y[0] {
			for k := range x[j] {
				ans[i][j] += x[i][k] * y[k][j]
			}
		}
	}

	return ans, nil
}
