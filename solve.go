package k3library

func Solve(l, u [][]float64, b []float64, idx []int) ([]float64, error) {

	b_ := make([]float64, len(b))

	for i, v := range idx {
		b_[i] = b[v]
	}

	v, err1 := ForwardSub(l, b_)
	if err1 != nil {
		return v, ErrCannotSolve
	}

	x, err2 := BackSub(u, v)
	if err2 != nil {
		return x, ErrCannotSolve
	}

	return x, nil
}
