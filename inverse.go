package k3library

func (m *Mat) Inverse() (Mat, error) {
	l, u, idx, _, err := m.LUDecomp()
	ans := NewMat(m.Col, m.Row)

	if err != nil {
		tmp := NewMat(0, 0)
		return tmp, ErrInvalid
	}

	e := makeIdentityMat(m.Col)

	for i := 0; i < m.Col; i++ {
		x, _ := Solve(l, u, e.M[i], idx)
		ans.M[i] = x
	}

	return ans.Transpose(), nil
}
