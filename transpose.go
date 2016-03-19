package k3library

func Transpose(a [][]float64) [][]float64 {
	ans := makeMat(len(a[0]), len(a))

	for i := range a[0] {
		for j := range a {
			ans[i][j] = a[j][i]
		}
	}

	return ans
}
