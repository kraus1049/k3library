package k3library

import (
	"errors"
)

func Euler(f FNCVec, t_i float64, x_i Vec, to, eps float64) (func(float64) (float64, error), error) {
	n := int((to - t_i) / eps)

	ans := make([]float64, n)

	t_i_ := t_i

	ans[0] = x_i.At(0)
	pre_tmp := x_i.Copy()
	for i := 1; i < n; i++ {
		if tmp, err := EulerNext(f, t_i_, pre_tmp, eps); err == nil {
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

func EulerNext(f FNCVec, t_i float64, x_i Vec, dt float64) (Vec, error) {
	if x_i_, err := f.Calc(t_i, x_i); err == nil {
		tmp, _ := Pro(x_i_, dt)
		ans, _ := Sum(x_i, tmp)
		return ans.(Vec), nil
	} else {
		return NewVec(0), err
	}
	return NewVec(0), errors.New("reaching here is bug")
}
