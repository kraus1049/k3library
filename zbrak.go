package k3library

import (
	"fmt"
	"os"
)

func Zbrak(start, end, num float64, div_n int, g func(float64) float64) (ans [][2]float64, ok bool) {

	defer func() {
		if x := recover(); x != nil {
			fmt.Fprintln(os.Stderr, x)
			ok = false
		} else {
			ok = true
		}
	}()

	if start > end {
		start, end = end, start
	} else if start == end || div_n <= 0 {
		panic("Zbrak: invalid argument")
	}

	f := func(x float64) float64 { return g(x) - num }

	var (
		x  float64 = start
		dx float64 = (end - start) / float64(div_n)
		fs float64 = f(x)
		fe float64 = 0
	)

	for i := 1; i <= div_n; i++ {
		x += dx
		fe = f(x)

		if fs*fe < 0 {
			ans = append(ans, [2]float64{x - dx, x})
			fs = fe
		} else if fs*fe == 0 {
			ans = append(ans, [2]float64{x - dx, x + dx})
			i++
			x += dx
			fs = f(x)
		} else {
			fs = fe
		}
	}

	return
}
