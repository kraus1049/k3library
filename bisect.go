package k3library

import (
	"fmt"
	"math"
)

func Bisect(start, end float64, g func(float64) float64, num, eps float64) (ans float64, err error) {

	defer func() {
		if x := recover(); x != nil {
			// fmt.Fprintln(os.Stderr, x)
			err = fmt.Errorf("%v",x)
		}
	}()

	f := func(x float64) float64 { return g(x) - num }

	fs := f(start)
	fe := f(end)

	if Epsequal(fs, 0, eps) {
		ans = start
		return
	} else if Epsequal(fe, 0, eps) {
		ans = end
		return
	}

	if fs*fe > 0 {
		panic("Bisect:Invalid argument")
	} else if start > end {
		start, end = end, start

		fs, fe = fe, fs
	}

	mid := (start + end) / 2.0
	fm := f(mid)
	cnt := (int)(math.Ceil(math.Log2((end - start)/eps)) * 2)

	for !Epsequal(fm,0,eps){
		if fs*fm < 0 {
			end = mid
		} else {
			start = mid
		}

		mid = (start + end) / 2.0
		fs = f(start)
		fm = f(mid)

		if cnt--;cnt <= 0{
			panic("Bisect:Infinite loop")
		}
	}
	ans = mid
	return
}
