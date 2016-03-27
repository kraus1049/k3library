package k3library

import (
	"math"
)

func Newton(start float64, f, g func(float64) float64, eps float64) (float64, error) {

	fs := f(start)

	if EpsEqual(fs, 0, eps) {
		return start, nil
	}

	gs := g(start)

	x_i := start
	var cnt int
	var x_in float64
	diff := math.MaxFloat64

	for !EpsEqual(fs, 0, eps) {
		if EpsEqual(gs, 0, 1e-3) {
			return -1, ErrInvalid
		}

		x_in = x_i - (fs / gs)

		diff_now := math.Abs(x_in - x_i)

		if EpsEqual(diff, diff_now, eps) {
			cnt++
		}

		if cnt > 10000 {
			return -1, ErrInfiniteLoop
		}

		diff = math.Abs(x_in - x_i)
		x_i = x_in
		fs = f(x_i)
		gs = g(x_i)

	}

	return x_in, nil
}
