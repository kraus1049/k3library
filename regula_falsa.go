package k3library

import (
	"fmt"
	"os"
)

func Regula_falsa(start, end float64, g func(float64) float64, num, eps float64) (ans float64, ok bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Fprintln(os.Stderr, x)
			ok = false
		} else {
			ok = true
		}
	}()

	f := func(x float64) float64 { return g(x) - num }
	var mid, pre_mid, fs, fe, fm float64

	fs, fe = f(start), f(end)
	mid = start - (fs * ((end - start) / (fe - fs)))
	fm = f(mid)

	if start == end {
		if Epsequal(fs, 0, eps) {
			ans = start
			return
		} else {
			panic("Regula_falsa: Invalid Argument")
		}
	} else if start > end {
		start, end = end, start
		fs, fe = fe, fs
	}

	if fs*fm > 0 {
		panic("Regula_falsa: Invalid Argument")
	}

	for !Epsequal(pre_mid, mid, eps) || !Epsequal(fm,0,eps) {
		if fs*fm < 0 {
			end = mid
			fe = f(end)
		} else {
			start = mid
			fs = f(start)
		}

		pre_mid, mid = mid, start-(fs*((end-start)/(fe-fs)))
		fm = f(mid)
	}

	ans = mid

	return

}
