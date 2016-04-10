package k3library

func Sum(xs ...interface{}) (interface{}, error) {
	if !allTypeEqual(xs...) {
		return nil, ErrInvalid
	}

	if len(xs) == 1 {
		if x, ok := xs[0].(float64); ok {
			return x, nil
		} else if v, ok := xs[0].(Vec); ok {
			return v, nil
		} else if m, ok := xs[0].(Mat); ok {
			return m, nil
		} else {
			return nil, ErrInvalid
		}

	} else if len(xs) >= 2 {
		switch xs[0].(type) {
		case float64:
			x, _ := xs[0].(float64)
			y, _ := xs[1].(float64)

			xs = append(xs[2:], x+y)
		case Vec:
			x, _ := xs[0].(Vec)
			y, _ := xs[1].(Vec)
			if x.Row != y.Row {
				return nil, ErrInvalid
			}

			plus := NewVec(x.Row)

			for i := 0; i < x.Row; i++ {
				tmp := x.At(i) + y.At(i)
				plus.Write(i, tmp)
			}

			xs = append(xs[2:], plus)

		case Mat:
			x, _ := xs[0].(Mat)
			y, _ := xs[1].(Mat)

			if x.Col != y.Col || x.Row != y.Row {
				return nil, ErrInvalid
			}

			plus := NewMat(x.Col, x.Row)

			for i := 0; i < x.Col; i++ {
				for j := 0; j < x.Row; j++ {
					tmp := x.At(i, j) + y.At(i, j)
					plus.Write(i, j, tmp)
				}
			}

			xs = append(xs[2:], plus)
		}

		ans, err := Sum(xs...)
		return ans, err

	}
	return nil, ErrInvalid
}

func Sub(x, y interface{}) (interface{}, error) {
	y_, err := Pro(-1.0, y)
	if err != nil {
		return nil, err
	}
	if i, ok := y_.(float64); ok {
		return Sum(x, i)
	} else if v, ok := y_.(Vec); ok {
		return Sum(x, v)
	} else if m, ok := y_.(Mat); ok {
		return Sum(x, m)
	}
	return nil, ErrInvalid

}

func Pro(xs ...interface{}) (interface{}, error) {
	if len(xs) == 1 {
		if x, ok := xs[0].(float64); ok {
			return x, nil
		} else if v, ok := xs[0].(Vec); ok {
			return v, nil
		} else if m, ok := xs[0].(Mat); ok {
			return m, nil
		} else {
			return nil, ErrInvalid
		}
	} else if len(xs) >= 2 {
		switch xs[0].(type) {
		case float64:
			x, _ := xs[0].(float64)

			switch xs[1].(type) {
			case float64:
				y, _ := xs[1].(float64)
				xs = append([]interface{}{x * y}, xs[2:]...)
			case Vec:
				y, _ := xs[1].(Vec)
				pro := y.Copy()
				for i := 0; i < pro.Row; i++ {
					tmp := pro.At(i) * x
					pro.Write(i, tmp)
				}
				xs = append([]interface{}{pro}, xs[2:]...)
			case Mat:
				y, _ := xs[1].(Mat)
				pro := y.Copy()
				for i := 0; i < pro.Col; i++ {
					for j := 0; j < pro.Row; j++ {
						tmp := pro.At(i, j) * x
						pro.Write(i, j, tmp)
					}
				}
				xs = append([]interface{}{pro}, xs[2:]...)
			default:
				return nil, ErrInvalid
			}
		case Vec:
			x, _ := xs[0].(Vec)

			switch xs[1].(type) {
			case float64:
				y, _ := xs[1].(float64)
				pro := x.Copy()
				for i := 0; i < pro.Row; i++ {
					tmp := pro.At(i) * y
					pro.Write(i, tmp)
				}
				xs = append([]interface{}{pro}, xs[2:]...)

			case Vec:
				y, _ := xs[1].(Vec)
				if x.Row != y.Row {
					return nil, ErrInvalid
				}

				sum := 0.0
				for i := 0; i < x.Row; i++ {
					sum += x.At(i) * y.At(i)
				}
				xs = append([]interface{}{sum}, xs[2:]...)
			case Mat:
				y, _ := xs[1].(Mat)
				if y.Col != 1 {
					return nil, ErrInvalid
				}

				m := NewMat(x.Row, y.Row)
				for i := 0; i < x.Row; i++ {
					for j := 0; j < y.Row; j++ {
						sum := 0.0
						for k := 0; k < y.Col; k++ {
							sum += x.At(i) * y.At(k, j)
						}
						m.Write(i, j, sum)
					}
				}
				xs = append([]interface{}{m}, xs[2:]...)
			default:
				return nil, ErrInvalid
			}
		case Mat:
			x, _ := xs[0].(Mat)
			switch xs[1].(type) {
			case float64:
				y, _ := xs[1].(float64)
				pro := x.Copy()
				for i := 0; i < pro.Col; i++ {
					for j := 0; j < pro.Row; j++ {
						tmp := pro.At(i, j) * y
						pro.Write(i, j, tmp)
					}
				}
				xs = append([]interface{}{pro}, xs[2:]...)

			case Vec:
				y, _ := xs[1].(Vec)

				if x.Row != y.Row {
					return nil, ErrInvalid
				}
				pro := NewVec(x.Col)
				for i := 0; i < x.Col; i++ {
					sum := 0.0
					for k := 0; k < y.Row; k++ {
						sum += x.At(i, k) * y.At(k)
					}
					pro.Write(i, sum)
				}
				xs = append([]interface{}{pro}, xs[2:]...)
			case Mat:
				y, _ := xs[1].(Mat)

				if x.Row != y.Col {
					return nil, ErrInvalid
				}
				pro := NewMat(x.Col, y.Row)

				for i := 0; i < x.Col; i++ {
					for j := 0; j < y.Row; j++ {
						sum := 0.0
						for k := 0; k < x.Row; k++ {
							sum += x.At(i, k) * y.At(k, j)
						}
						pro.Write(i, j, sum)
					}
				}
				xs = append([]interface{}{pro}, xs[2:]...)
			default:
				return nil, ErrInvalid
			}

		default:
			return nil, ErrInvalid
		}

		ans, err := Pro(xs...)
		return ans, err
	}

	return nil, ErrInvalid

}
