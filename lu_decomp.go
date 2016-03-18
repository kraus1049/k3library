package k3library

func LUDecomp(a [][]float64) ([][]float64, [][]float64, error) {

	if !IsSquareMat(a) {
		return a, a, ErrInvalid
	}

	l := make([][]float64, len(a))
	for i := range l {
		l[i] = make([]float64, len(a[i]))
	}

	for i := range l {
		l[i][i] = 1
	}

	u := make([][]float64, len(a))
	for i := range u {
		u[i] = make([]float64, len(a[i]))
	}

	for i := range a {
		for j := i; j < len(a[i]); j++ {
			sgm := 0.0

			for k := 0; k < i; k++ {
				sgm += l[i][k] * u[k][j]
			}

			u[i][j] = a[i][j] - sgm
		}

		if u[i][i] == 0 {
			return l, u, ErrCannotSolve
		}

		for j := i + 1; j < len(a); j++ {
			sgm := 0.0

			for k := 0; k < i; k++ {
				sgm += l[j][k] * u[k][i]
			}
			l[j][i] = (a[j][i] - sgm) / u[i][i]
		}

	}

	return l, u, nil
}
