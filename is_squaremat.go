package k3library

func IsSquareMat(mat [][]float64) bool {
	for i := range mat {
		if len(mat) != len(mat[i]) {
			return false
		}
	}
	return true
}
