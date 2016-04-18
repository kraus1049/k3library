package k3library

func (a *Mat) Transpose() Mat {
	ans := NewMat(a.row, a.col)

	for i := 0; i < a.row; i++ {
		for j := 0; j < a.col; j++ {
			ans.Write(i, j, a.At(j, i))
		}
	}

	return ans
}
