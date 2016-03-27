package k3library

func Secant(start, end float64, g func(float64) float64, num, eps float64) (float64, error) {

	f := func(x float64) float64 { return g(x) - num }
	var fs, fe float64

	fs, fe = f(start), f(end)

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

	for !EpsEqual(start, end, eps) {
		start, end = start-(fs*((end-start)/(fe-fs))), start
		fs, fe = f(start), f(end)
	}

	return end, nil
}
