package k3library

func (mat *Mat) IsSquareMat() bool {
	for i := 0; i < mat.Col; i++ {
		if mat.Col != len(mat.M[i].V) {
			return false
		}
	}
	return true
}
