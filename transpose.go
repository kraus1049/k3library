package k3library

func (a *Mat) Transpose() Mat {
	ans := NewMat(a.Row, a.Col)

	for i := 0; i < a.Row; i++ {
		for j := 0; j < a.Col; j++ {
			ans.Write(i, j, a.At(j, i))
		}
	}

	return ans
}
