package k3library

func (mat *Mat) IsSquareMat() bool {
	for i := 0; i < mat.col; i++ {
		if mat.col != len(mat.m[i].v) {
			return false
		}
	}
	return true
}
