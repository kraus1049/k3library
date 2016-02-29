package k3library

func Csv2func(filepath string, parser func(string) (map[string]string, [][2]float64, error)) (f func(float64) (float64, error), err error) {

	_, xys, err := parser(filepath)
	if err != nil {
		return nil, err
	}

	x_min, x_max := xys[0][0], xys[len(xys)-1][0]

	f = func(x float64) (float64, error) {
		if x < x_min || x_max < x {
			return 0, ErrOutOfRange
		} else if x == x_max {
			return xys[len(xys)-1][1], nil
		} else if x == x_min {
			return xys[0][1], nil
		}

		var i int
		for i = range xys {
			if x < xys[i][0] {
				return xys[i][1], nil
			}
		}
		return 0, ErrInvalid
	}

	return f, nil
}
