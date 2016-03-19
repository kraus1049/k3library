package k3library

func Transpose(a [][]float64) [][]float64 {
	ans := make([][]float64, len(a[0]))
	for i := range ans {
		ans[i] = make([]float64, len(a))
	}

	for i := range a[0] {
		for j := range a {
			ans[i][j] = a[j][i]
		}
	}

	return ans
}
