package k3library

func Solve(l, u Mat, b Vec, idx []int) (Vec, error) {

	b_ := NewVec(b.Row)

	for i, v := range idx {
		b_.Write(i, b.At(v))
	}

	v, err1 := ForwardSub(l, b_)
	if err1 != nil {
		tmp := NewVec(0)
		return tmp, ErrCannotSolve
	}

	x, err2 := BackSub(u, v)
	if err2 != nil {
		tmp := NewVec(0)
		return tmp, ErrCannotSolve
	}

	return x, nil
}

func ProveSolve(a Mat, b Vec, eps float64) (Vec, error) {
	l, u, idx, _, err := a.LUDecomp()

	if err != nil {
		tmp := NewVec(0)
		return tmp, ErrCannotSolve
	}

	x, err := Solve(l, u, b, idx)

	if err != nil {
		tmp := NewVec(0)
		return tmp, ErrCannotSolve
	}

	ax, err := Pro(a, x)
	if err != nil {
		tmp := NewVec(0)
		return tmp, ErrCannotSolve
	}

	b_, err := Sub(ax, b)
	if err != nil {
		tmp := NewVec(0)
		return tmp, ErrCannotSolve
	}

	dx, err := Solve(l, u, b_.(Vec), idx)
	if err != nil {
		tmp := NewVec(0)
		return tmp, ErrCannotSolve
	}
	ans, err := Sub(x, dx)

	if err != nil {
		tmp := NewVec(0)
		return tmp, ErrCannotSolve
	}

	zero := NewVec(b.Row)

	for cnt := 10; cnt >= 0; cnt-- {
		if VecEpsEqual(dx, zero, eps) {
			break
		}

		ax, err := Pro(a, dx)
		if err != nil {
			tmp := NewVec(0)
			return tmp, ErrCannotSolve
		}

		b_, err = Sub(ax, b)
		if err != nil {
			tmp := NewVec(0)
			return tmp, ErrCannotSolve
		}

		dx, err = Solve(l, u, b_.(Vec), idx)
		if err != nil {
			tmp := NewVec(0)
			return tmp, ErrCannotSolve
		}

		ans, err = Sub(x, dx)
		if err != nil {
			tmp := NewVec(0)
			return tmp, ErrCannotSolve
		}

	}

	return ans.(Vec), nil
}
