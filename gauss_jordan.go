package k3library

func GaussJordan(a [][]float64, b []float64) (x []float64, err error) {
	if !canSimultaneousEquSolve(a, b) {
		err = ErrInvalid
		return
	}

	idx := serialNum(len(a))

	a_ := MatCopy(a)
	b_ := VecCopy(b)

	if e := forwardDelIdx(a_, b_, idx); e != nil {
		err = e
		return
	}

	x, _ = backSubIdx(a_, b_, idx)

	return
}
