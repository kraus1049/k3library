package k3library

func GaussJordan(a [][]float64, b []float64) ([]float64, error) {
	if !canSimultaneousEquSolve(a, b) {
		return nil, ErrInvalid
	}

	idx := serialNum(len(a))

	a_ := MatCopy(a)
	b_ := VecCopy(b)

	if err := forwardDelIdx(a_, b_, idx); err != nil {
		return nil, err
	}

	x, _ := backSubIdx(a_, b_, idx)

	return x, nil
}
