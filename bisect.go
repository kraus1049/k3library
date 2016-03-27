package k3library

import (
	"math"
)

func Bisect(start, end float64, g func(float64) float64, num, eps float64) (float64, error) {

	f := func(x float64) float64 { return g(x) - num }

	fs := f(start)
	fe := f(end)

	if EpsEqual(fs, 0, eps) {
		return start, nil
	} else if EpsEqual(fe, 0, eps) {
		return end, nil
	}

	if fs*fe > 0 {
		return -1, ErrInvalid
	} else if start > end {
		start, end = end, start
		fs, fe = fe, fs
	}

	mid := (start + end) / 2.0
	fm := f(mid)
	cnt := (int)(math.Ceil(math.Log2((end-start)/eps)) * 2)

	for !EpsEqual(fm, 0, eps) {
		if fs*fm < 0 {
			end = mid
		} else {
			start = mid
		}

		mid = (start + end) / 2.0
		fs = f(start)
		fm = f(mid)

		if cnt--; cnt <= 0 {
			return -1, ErrInfiniteLoop
		}
	}
	return mid, nil
}
