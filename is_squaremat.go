package k3library

func IsSquareMat(mat Mat) bool {
	for i := 0; i < mat.Col; i++ {
		if mat.Col != len(mat.M[i].V) {
			return false
		}
	}
	return true
}
