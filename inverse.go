package k3library

func Inverse(a [][]float64) ([][]float64, error) {
	l, u, idx, _, err := LUDecomp(a)
	ans := makeMat(len(a), len(a[0]))

	if err != nil {
		return nil, ErrInvalid
	}

	e := makeIdentityMat(len(a))

	for i := range a {
		x, _ := Solve(l, u, e[i], idx)
		ans[i] = x
	}

	return Transpose(ans), nil
}
