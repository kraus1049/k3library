package k3library

func Det(u [][]float64, sgn int) float64 {
	det := 1.0

	for i := range u {
		det *= u[i][i]
	}

	return det
}
