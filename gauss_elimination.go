package k3library

func GaussElimination(a Mat, b Vec) (Vec, error) {
	if !canSimultaneousEquSolve(a, b) {
		return Vec{[]float64{}, 0}, ErrInvalid
	}

	idx := serialNum(a.col)

	a_ := a.Copy()
	x_ := b.Copy()
	x := NewVec(x_.row)

	for i := 0; i < a_.col; i++ {

		if a_.At(idx[i], i) == 0 {
			if i == a_.col-1 {
				return Vec{[]float64{}, 0}, ErrCannotSolve
			}

			tmp := make([]float64, 0)
			for j := i + 1; j < a_.col; j++ {
				maxnum, _ := max(a_.m[idx[j]])
				tmp = append(tmp, a_.At(idx[j], i)/maxnum)
			}
			tmp_ := NewVecSet(tmp...)

			if maxnum, maxidx := max(tmp_); maxnum != 0 {
				idx[i], idx[maxidx+i+1] = idx[maxidx+i+1], idx[i]
			} else {
				return Vec{[]float64{}, 0}, ErrCannotSolve
			}
		}

		num := a_.At(idx[i], i)

		for j := i; j < a_.col; j++ {
			tmp := a_.At(idx[i], j) / num
			a_.Write(idx[i], j, tmp)
		}

		x_.Write(idx[i], x_.At(idx[i])/num)

		for j := 0; j < a_.col; j++ {
			if idx[i] == idx[j] {
				continue
			}

			per := a_.At(idx[j], i) / a_.At(idx[i], i)

			for k := i; k < a_.col; k++ {
				a_.Write(idx[j], k, a_.At(idx[j], k)-per*a_.At(idx[i], k))
			}

			x_.Write(idx[j], x_.At(idx[j])-per*x_.At(idx[i]))
		}

	}

	for i, changed := range idx {
		x.Write(changed, x_.At(i))
	}

	return x, nil
}
