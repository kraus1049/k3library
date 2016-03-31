package k3library

func Det(u Mat, sgn int) float64 {
	det := 1.0

	for i := 0; i < u.Col; i++ {
		det *= u.At(i, i)
	}

	return det * float64(sgn)
}
