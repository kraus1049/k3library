package k3library

func RungeKutta(f FNCVec, t_i float64, x_i Vec, to, eps float64) (func(float64) (float64, error), error) {
	n := int((to - t_i) / eps)

	ans := make([]float64, n)

	t_i_ := t_i

	ans[0] = x_i.At(0)
	pre_tmp := x_i.Copy()
	for i := 1; i < n; i++ {
		if tmp, err := RungeKuttaNext(f, t_i_, pre_tmp, eps); err == nil {
			t_i_ += eps
			ans[i] = tmp.At(0)
			pre_tmp = tmp

		} else {
			return nil, err
		}
	}

	g := func(t float64) (float64, error) {
		if t < t_i || t_i_ < t {
			return -1, ErrOutOfRange
		}

		idx := int((t - t_i) / eps)
		return ans[idx], nil
	}
	return g, nil
}

func RungeKuttaNext(f FNCVec, t_i float64, x_i Vec, dt float64) (Vec, error) {

	k1, err := f.Calc(t_i, x_i)

	if err != nil {
		return NewVec(0), err
	}

	dt_2 := dt / 2.0
	t_i_dt_2 := t_i + dt_2

	tmp, _ := Pro(k1, dt_2)
	tmp, _ = Sum(x_i, tmp)

	k2, err := f.Calc(t_i_dt_2, tmp.(Vec))
	k2_2, _ := Pro(2.0, k2)

	if err != nil {
		return NewVec(0), err
	}

	tmp, _ = Pro(k2, dt_2)
	tmp, _ = Sum(x_i, tmp)

	k3, err := f.Calc(t_i_dt_2, tmp.(Vec))
	k3_2, _ := Pro(2.0, k3)

	if err != nil {
		return NewVec(0), err
	}

	tmp, _ = Pro(k3, dt_2)
	tmp, _ = Sum(x_i, tmp)

	k4, err := f.Calc(t_i+dt, tmp.(Vec))

	if err != nil {
		return NewVec(0), err
	}

	tmp, _ = Sum(k1, k2_2, k3_2, k4)
	tmp, _ = Pro(tmp, 1.0/6.0, dt)
	tmp, _ = Sum(x_i, tmp)

	return tmp.(Vec), nil
}
