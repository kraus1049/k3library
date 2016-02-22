package k3library

func Secant(start, end float64, g func(float64) float64, num, eps float64) (ans float64, err error) {
	// defer func() {
	// 	if x := recover(); x != nil {
	// 		err = fmt.Errorf("%v", x)
	// 	}
	// }()

	f := func(x float64) float64 { return g(x) - num }
	var fs, fe float64

	fs, fe = f(start), f(end)

	if start == end {
		if Epsequal(fs, 0, eps) {
			ans = start
			return
		} else {
			panic("Secant:Invalid argument")
			err = ErrInvalid
			return
		}
	} else if start > end {
		start, end = end, start
		fs, fe = fe, fs
	}

	for !Epsequal(start, end, eps) {
		start, end = start-(fs*((end-start)/(fe-fs))), start
		fs, fe = f(start), f(end)
	}

	ans = end
	return
}
