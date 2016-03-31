package k3library

func BackSub(a Mat, b Vec) (Vec, error) {
	x := NewVec(b.Row)

	for i := a.Col - 1; i >= 0; i-- {
		if a.At(i, i) == 0 {
			return Vec{[]float64{}, 0}, ErrCannotSolve
		}
		sgm := 0.0
		for j := a.Col - 1; j > i; j-- {
			sgm += a.At(i, j) * x.At(j)
		}
		x.Write(i, (b.At(i)-sgm)/a.At(i, i))
	}

	return x, nil
}
