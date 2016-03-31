package k3library

func Inverse(a Mat) (Mat, error) {
	l, u, idx, _, err := LUDecomp(a)
	ans := NewMat(a.Col, a.Row)

	if err != nil {
		tmp := NewMat(0, 0)
		return tmp, ErrInvalid
	}

	e := makeIdentityMat(a.Col)

	for i := 0; i < a.Col; i++ {
		x, _ := Solve(l, u, e.M[i], idx)
		ans.M[i] = x
	}

	return Transpose(ans), nil
}
