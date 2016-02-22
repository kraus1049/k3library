package k3library

func Sgn(x float64) (ans float64, ok bool) {
	defer func() {
		if x := recover(); x != nil {
			ok = false
		} else {
			ok = true
		}
	}()

	if x > 0 {
		ans = 1
	} else if x < 0 {
		ans = -1
	} else {
		ans = 0
	}

	return
}