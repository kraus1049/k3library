package k3library

func Inverse(a [][]float64) ([][]float64, error) {

	l, u, idx, _, err := LUDecomp(a)
	ans := make([][]float64, len(a))
	for i := range ans {
		ans[i] = make([]float64, len(a[i]))
	}

	if err != nil {
		return l, ErrInvalid
	}

	for i := range a {

		e := make([]float64, len(a))
		e[i] = 1

		x, err2 := Solve(l, u, e, idx)

		if err2 != nil {
			return l, ErrInvalid
		}

		ans[i] = x

	}

	return Transpose(ans), nil
}
