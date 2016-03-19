package k3library

func LUDecomp(a [][]float64) ([][]float64, [][]float64, []int, int, error) {

	idx := serialNum(len(a))

	if !IsSquareMat(a) {
		return a, a, idx, 1, ErrInvalid
	}

	var sgn int = 1

	l := make([][]float64, len(a))
	for i := range l {
		l[i] = make([]float64, len(a[i]))
	}

	u := make([][]float64, len(a))
	for i := range u {
		u[i] = make([]float64, len(a[i]))
	}

	for i := range a {
		flag := false
		for h := i; h < len(a); h++ {
			for j := i; j < len(a[i]); j++ {
				sgm := 0.0

				for k := 0; k < i; k++ {
					sgm += l[idx[i]][k] * u[k][j]
				}

				u[i][j] = a[idx[i]][j] - sgm
			}

			if u[i][i] == 0 && h+1 < len(a) {
				idx[i], idx[h+1] = idx[h+1], idx[i]
				flag = true
			} else {
				break
			}

		}

		if flag {
			sgn = -sgn
		}

		l[idx[i]][i] = 1

		for j := i + 1; j < len(a); j++ {
			sgm := 0.0

			for k := 0; k < i; k++ {
				sgm += l[idx[j]][k] * u[k][i]
			}
			l[idx[j]][i] = (a[idx[j]][i] - sgm) / u[i][i]

		}

	}

	return swapMatIdx(l, idx), u, idx, sgn, nil
}
