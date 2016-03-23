package k3library

func Trapezoid(f func(float64) (float64, error), a, b, eps float64) (float64, error) {
	if a >= b {
		return -1, ErrInvalid
	}

	n := 1
	pre_ans, err := trapezoid1(f, a, b, &n)

	if err != nil {
		return -1, err
	}

	ans, err := trapezoid(f, a, b, &n, pre_ans)

	if err != nil {
		return -1, err
	}

	cnt := 0

	for !EpsEqual(pre_ans, ans, eps) {
		pre_ans = ans
		ans, err = trapezoid(f, a, b, &n, pre_ans)

		if err != nil {
			return -1, err
		}

		if cnt++; cnt >= 100 {
			break
		}
	}

	return ans, nil
}

func trapezoid1(f func(float64) (float64, error), a, b float64, n *int) (float64, error) {
	h := float64(b-a) / float64(*n)

	sum := 0.0

	for i := 1; i < *n; i++ {
		if tmp, err := f(a + float64(i)*h); err == nil {
			sum += tmp
		} else {
			return -1, err
		}
	}

	*n *= 2
	return sum * h / 2.0, nil
}

func trapezoid(f func(float64) (float64, error), a, b float64, n *int, pre_ans float64) (float64, error) {
	*n *= 2
	h := float64(b-a) / float64(*n)

	sum := 0.0
	for i := 1; i < *n; i += 2 {
		if tmp, err := f(a + float64(i)*h); err == nil {
			sum += tmp
		} else {
			return -1, err
		}
	}

	return pre_ans/2.0 + h*sum, nil

}
