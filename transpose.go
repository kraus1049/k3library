package k3library

func Transpose(a Mat) Mat {
	ans := NewMat(a.Row, a.Col)

	for i := 0; i < a.Row; i++ {
		for j := 0; j < a.Col; j++ {
			ans.Write(i, j, a.At(j, i))
		}
	}

	return ans
}
