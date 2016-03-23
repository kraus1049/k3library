package k3library

func Rectangle(f func(float64) float64, a, b, eps float64) (float64, error) {
	if a > b {
		return -1, ErrInvalid
	}

	n := 1
	pre_ans := rectangle1(f, a, b, &n)
	ans := rectangle(f, a, b, &n, pre_ans)

	cnt := 0

	for !EpsEqual(pre_ans, ans, eps) {

		pre_ans = ans
		ans = rectangle(f, a, b, &n, pre_ans)

		cnt++
		if cnt >= 100 {
			break
		}
	}

	return ans, nil
}

func rectangle1(f func(float64) float64, a, b float64, n *int) float64 {
	h := float64(b-a) / float64(*n)

	sum := 0.0

	for i := 0; i < *n; i++ {
		sum += f(a + float64(i)*h)
	}

	*n *= 2
	return h * sum
}

func rectangle(f func(float64) float64, a, b float64, n *int, pre_ans float64) float64 {
	*n *= 2
	h := float64(b-a) / float64(*n)

	sum := 0.0

	for i := 1; i < *n; i += 2 {
		sum += f(a + float64(i)*h)
	}

	return pre_ans/2.0 + h*sum
}
