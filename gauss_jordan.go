package k3library

func GaussJordan(a Mat, b Vec) (Vec, error) {
	if !canSimultaneousEquSolve(a, b) {
		return Vec{[]float64{}, 0}, ErrInvalid
	}

	idx := serialNum(a.Col)

	a_ := a.Copy()
	b_ := b.Copy()

	if err := forwardDelIdx(a_, b_, idx); err != nil {
		return Vec{[]float64{}, 0}, err
	}

	x, _ := backSubIdx(a_, b_, idx)

	return x, nil
}
