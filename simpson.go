package k3library

func Simpson(f func(float64) float64, a, b, eps float64) (float64, error) {

	if a >= b {
		return -1, ErrInvalid
	}

	simp := func(pre_ans, ans float64) float64 {
		return (4*ans - pre_ans) / 3.0
	}

	n := 1
	t_pre_ans := trapezoid1(f, a, b, &n)
	t_ans := trapezoid(f, a, b, &n, t_pre_ans)

	s_pre_ans := simp(t_pre_ans, t_ans)
	t_pre_ans = t_ans
	t_ans = trapezoid(f, a, b, &n, t_pre_ans)
	s_ans := simp(t_pre_ans, t_ans)

	for !EpsEqual(s_pre_ans, s_ans, eps) {
		s_pre_ans = s_ans
		t_pre_ans = t_ans
		t_ans = trapezoid(f, a, b, &n, t_pre_ans)
		s_ans = simp(t_pre_ans, t_ans)

	}

	return s_ans, nil
}
