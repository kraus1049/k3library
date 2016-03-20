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

func ProveSolve(a [][]float64, b []float64, eps float64) ([]float64, error) {
	l, u, idx, _, err := LUDecomp(a)

	if err != nil {
		return nil, ErrCannotSolve
	}

	x, err := Solve(l, u, b, idx)

	if err != nil {
		return nil, ErrCannotSolve
	}

	ax, err := MPro(a, Transpose([][]float64{x}))
	if err != nil {
		return nil, ErrCannotSolve
	}
	ax = Transpose(ax)
	b_, err := VSub(ax[0], b)
	if err != nil {
		return nil, ErrCannotSolve
	}

	dx, err := Solve(l, u, b_, idx)
	if err != nil {
		return nil, ErrCannotSolve
	}
	ans, err := VSub(x, dx)

	if err != nil {
		return nil, ErrCannotSolve
	}

	zero := make([]float64, len(b))

	for cnt := 10; cnt >= 0; cnt-- {
		if SliceEpsEqual(dx, zero, eps) {
			break
		}

		ax, err = MPro(a, Transpose([][]float64{dx}))
		if err != nil {
			return nil, ErrCannotSolve
		}

		ax = Transpose(ax)
		b_, err = VSub(ax[0], b)
		if err != nil {
			return nil, ErrCannotSolve
		}

		dx, err = Solve(l, u, b_, idx)
		if err != nil {
			return nil, ErrCannotSolve
		}

		ans, err = VSub(x, dx)
		if err != nil {
			return nil, ErrCannotSolve
		}

	}

	return ans, nil
}
