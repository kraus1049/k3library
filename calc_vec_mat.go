package k3library

func MPro(x, y [][]float64) ([][]float64, error) {

	if len(x[0]) != len(y) {
		return nil, ErrInvalid
	}

	ans := makeMat(len(x), len(y[0]))

	for i := range x {
		for j := range y[0] {
			for k := range x[j] {
				ans[i][j] += x[i][k] * y[k][j]
			}
		}
	}

	return ans, nil
}
