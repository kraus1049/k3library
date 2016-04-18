package k3library

func (a *Mat) LUDecomp() (Mat, Mat, []int, int, error) {

	idx := serialNum(a.col)

	if !a.IsSquareMat() {
		tmp := NewMat(0, 0)
		return tmp, tmp, nil, -1, ErrInvalid
	}

	var sgn int = 1

	l := NewMat(a.col, a.row)
	u := NewMat(a.col, a.row)

	for i := 0; i < a.col; i++ {
		flag := false
		for h := i; h < a.col; h++ {
			for j := i; j < a.row; j++ {
				sgm := 0.0

				for k := 0; k < i; k++ {
					sgm += l.At(idx[i], k) * u.At(k, j)
				}

				u.Write(i, j, a.At(idx[i], j)-sgm)
			}

			if u.At(i, i) == 0 && h+1 < a.col {
				idx[i], idx[h+1] = idx[h+1], idx[i]
				flag = true
			} else {
				break
			}

		}

		if flag {
			sgn = -sgn
		}

		l.Write(idx[i], i, 1)

		for j := i + 1; j < a.col; j++ {
			sgm := 0.0

			for k := 0; k < i; k++ {
				sgm += l.At(idx[j], k) * u.At(k, i)
			}
			l.Write(idx[j], i, (a.At(idx[j], i)-sgm)/u.At(i, i))

		}

	}

	return swapMatIdx(l, idx), u, idx, sgn, nil
}
