package k3library

import (
	"math"
)

func Brent(start, end float64, g func(float64) float64, num, eps float64) (ans float64, err error) {

	f := func(x float64) float64 { return g(x) - num }
	const N int = 1000

	var (
		a                     float64 = start
		b                     float64 = end
		c                     float64 = end
		fa                    float64 = f(a)
		fb                    float64 = f(b)
		fc                    float64 = fb
		d, e, min1, min2, min float64
		p, q, r, s, eps1, xm  float64
	)

	if fa*fb > 0 {
		err = ErrInvalid
		return
	}

	for i := 1; i <= N; i++ {
		if fb*fc > 0 {
			c, fc = a, fa
			d = b - a
			e = d
		}

		if math.Abs(fc) < math.Abs(fb) {
			a, fa = b, fb
			b, fb = c, fc
			c, fc = a, fa
		}

		eps1 = 2*eps*math.Abs(b) + 0.5*eps
		xm = (c - b) / 2

		if math.Abs(xm) < eps1 || Epsequal(fb, 0, eps) {
			ans = b
			return
		}

		if math.Abs(e) >= eps1 && math.Abs(fa) > math.Abs(fb) {
			s = fb / fa
			if a == c {
				p = 2 * xm * s
				q = 1 - p
			} else {
				q = fa / fc
				r = fb / fc
				p = s * (2*xm*q*(q-r) - (b-a)*(r-1))
				q = (q - 1) * (r - 1) * (s - 1)
			}

			if p > 0 {
				q = -q
			}

			p = math.Abs(p)
			min1 = 3*xm*q - math.Abs(eps1*q)
			min2 = math.Abs(e * q)

			if min1 > min2 {
				min = min2
			} else {
				min = min1
			}

			if 2*p < min {
				e, d = d, p/q
			} else {
				d, e = xm, d
			}
		} else {
			d, e = xm, d
		}
		a, fa = b, fb
		if math.Abs(d) > eps1 {
			b += d
		} else {
			if xm >= 0 {
				b += math.Abs(eps1)
			} else {
				b -= math.Abs(eps1)
			}

		}
		fb = f(b)
	}

	return
}
