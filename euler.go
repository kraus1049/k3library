package k3library

import (
	"errors"
)

func Euler(f func(float64) (float64, error), t_i, x_i, to, eps float64) (func(float64) (float64, error), error) {
	n := int((to - t_i) / eps)

	ans := make([]float64, n)
	err := errors.New("")

	t_i_ := t_i

	ans[0] = x_i
	for i := 1; i < n; i++ {
		if ans[i], err = euler(f, t_i_, ans[i-1], eps); err == nil {
			t_i_ += eps
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

func euler(f func(float64) (float64, error), t_i, x_i, dt float64) (float64, error) {
	if x_i_, err := f(t_i); err == nil {
		return x_i + x_i_*dt, nil
	} else {
		return -1, err
	}
	return -1, errors.New("reaching here is bug")
}
