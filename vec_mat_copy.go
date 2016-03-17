package k3library

func VecCopy(vec []float64) []float64 {
	vec2 := make([]float64, len(vec))
	copy(vec2, vec)
	return vec2
}

func MatCopy(mat [][]float64) [][]float64 {

	mat2 := make([][]float64, len(mat))

	for i := range mat {
		mat2[i] = VecCopy(mat[i])
	}

	return mat2
}
