package k3library

func (m *Mat) Inverse() (Mat, error) {
	l, u, idx, _, err := m.LUDecomp()
	ans := NewMat(m.col, m.row)

	if err != nil {
		tmp := NewMat(0, 0)
		return tmp, ErrInvalid
	}

	e := makeIdentityMat(m.col)

	for i := 0; i < m.col; i++ {
		x, err := Solve(l, u, e.m[i], idx)
		if err != nil {
			return NewMat(0, 0), ErrInvalid
		}

		ans.m[i] = x
	}

	return ans.Transpose(), nil
}
