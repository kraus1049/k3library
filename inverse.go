package k3library

func Inverse(a [][]float64) ([][]float64, error) {

	l, u, idx, _, err := LUDecomp(a)
	ans := makeMat(len(a), len(a[0]))

	if err != nil {
		return l, ErrInvalid
	}

	e := makeIdentityMat(len(a))

	for i := range a {
		x, err2 := Solve(l, u, e[i], idx)

		if err2 != nil {
			return l, ErrInvalid
		}

		ans[i] = x
	}

	return Transpose(ans), nil
}
