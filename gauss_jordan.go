package k3library

func GaussJordan(a [][]float64, b []float64) (x []float64, err error) {
	if !IsSquareMat(a) || (len(a) != len(b)) {
		err = ErrInvalid
		return
	}

	idx := serialNum(len(a))

	a_ := MatCopy(a)
	b_ := VecCopy(b)

	for i := 0; i < len(a_); i++ {

		if a_[idx[i]][i] == 0 {
			if i == len(a_)-1 {
				err = ErrCannotSolve
				return
			}

			tmp := make([]float64, 0)
			for j := i + 1; j < len(a_); j++ {
				tmp = append(tmp, a_[idx[j]][i])
			}

			if maxnum, maxidx := max(tmp); maxnum != 0 {
				idx[i], idx[maxidx+i+1] = idx[maxidx+i+1], idx[i]
			} else {
				err = ErrCannotSolve
				return
			}
		}

		for j := i + 1; j < len(a_); j++ {
			per := a_[idx[j]][i] / a_[idx[i]][i]
			for k := i; k < len(a_); k++ {
				a_[idx[j]][k] -= per * a_[idx[i]][k]
			}

			b_[idx[j]] -= per * b_[idx[i]]
		}
	}

	x, _ = backSubIdx(a_, b_, idx)

	return
}
