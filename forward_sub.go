package k3library

func ForwardSub(a [][]float64, b []float64) ([]float64, error) {
	x := make([]float64, len(b))

	for i := range a {
		if a[i][i] == 0 {
			return x, ErrCannotSolve
		}
		sgm := 0.0
		for j := 0; j < i; j++ {
			sgm += a[i][j] * x[j]
		}
		x[i] = (b[i] - sgm) / a[i][i]
	}

	return x, nil
}
