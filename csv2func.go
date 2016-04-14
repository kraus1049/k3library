package k3library

func Csv2func(filepath string, parser func(string) (map[string]string, [][2]float64, error)) (func(float64) (float64, error), float64, float64, error) {

	_, xys, err := parser(filepath)
	if err != nil {
		return nil, -1, -1, err
	}

	xys = nub(xys)

	x_min, x_max := xys[0][0], xys[len(xys)-1][0]

	f := func(x float64) (float64, error) {
		if x < x_min || x_max < x {
			return 0, ErrOutOfRange
		} else if x == x_max {
			return xys[len(xys)-1][1], nil
		} else if x == x_min {
			return xys[0][1], nil
		}

		return binSearch(xys, x)
	}

	return f, x_min, x_max, nil
}

func binSearch(xys [][2]float64, x float64) (float64, error) {
	low, high := 0, len(xys)-1

	var ix int
	for low <= high {
		ix = (low + high) / 2

		if x == xys[ix][0] {
			return xys[ix][1], nil
		} else if x < xys[ix][0] {
			high = ix - 1
		} else {
			low = ix + 1
		}
	}
	return xys[ix-1][1], nil
}

func nub(x [][2]float64) [][2]float64 {
	ans := make([][2]float64, 0)
	ans = append(ans, x[0])

	for i := 1; i < len(x)-2; i++ {
		if x[i][1] != x[i-1][1] {
			ans = append(ans, x[i])
		}
	}

	ans = append(ans, x[len(x)-1])

	return ans
}
