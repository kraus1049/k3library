package k3library

import (
	"fmt"
	"math"
	"os"
)

func Newton(start float64, f, g func(float64) float64, eps float64) (ans float64, ok bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Fprintln(os.Stderr, x)
			ok = false
		} else {
			ok = true
		}
	}()

	fs := f(start)

	if Epsequal(fs, 0, eps) {
		ans = start
		return
	}

	gs := g(start)

	x_i := start
	var cnt int
	var x_in float64
	diff := math.MaxFloat64

	for !Epsequal(fs, 0, eps) {
		if Epsequal(gs, 0, 1e-3) {
			panic("Newton:invalid argument")
		}

		x_in = x_i - (fs / gs)

		diff_now := math.Abs(x_in - x_i)

		if Epsequal(diff, diff_now, eps) {
			cnt++
		}

		if cnt > 10000 {
			panic("Newton:loop detection")
		}

		diff = math.Abs(x_in - x_i)
		x_i = x_in
		fs = f(x_i)
		gs = g(x_i)

	}
	ans = x_in

	return
}
