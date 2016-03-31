package k3library

func RegulaFalsa(start, end float64, g func(float64) float64, num, eps float64) (float64, error) {

	f := func(x float64) float64 { return g(x) - num }
	var mid, pre_mid, fs, fe, fm float64

	fs, fe = f(start), f(end)
	mid = start - (fs * ((end - start) / (fe - fs)))
	fm = f(mid)

	if start == end {
		if EpsEqual(fs, 0, eps) {
			return start, nil
		} else {
			return -1, ErrInvalid
		}
	} else if start > end {
		start, end = end, start
		fs, fe = fe, fs
	}

	if fs*fm > 0 {
		return -1, ErrInvalid
	}

	for !EpsEqual(pre_mid, mid, eps) || !EpsEqual(fm, 0, eps) {
		if fs*fm < 0 {
			end = mid
			fe = f(end)
		} else {
			start = mid
			fs = f(start)
		}

		pre_mid, mid = mid, start-(fs*((end-start)/(fe-fs)))
		fm = f(mid)
	}

	return mid, nil

}
