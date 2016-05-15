package k3library

import (
	"math"
)

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
			if x.row != y.row {
				return nil, ErrInvalid
			}

			plus := NewVec(x.row)

			for i := 0; i < x.row; i++ {
				tmp := x.At(i) + y.At(i)
				plus.Write(i, tmp)
			}

			xs = append(xs[2:], plus)

		case Mat:
			x, _ := xs[0].(Mat)
			y, _ := xs[1].(Mat)

			if x.col != y.col || x.row != y.row {
				return nil, ErrInvalid
			}

			plus := NewMat(x.col, x.row)

			for i := 0; i < x.col; i++ {
				for j := 0; j < x.row; j++ {
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
				for i := 0; i < pro.row; i++ {
					tmp := pro.At(i) * x
					pro.Write(i, tmp)
				}
				xs = append([]interface{}{pro}, xs[2:]...)
			case Mat:
				y, _ := xs[1].(Mat)
				pro := y.Copy()
				for i := 0; i < pro.col; i++ {
					for j := 0; j < pro.row; j++ {
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
				for i := 0; i < pro.row; i++ {
					tmp := pro.At(i) * y
					pro.Write(i, tmp)
				}
				xs = append([]interface{}{pro}, xs[2:]...)

			case Vec:
				y, _ := xs[1].(Vec)
				if x.row != y.row {
					return nil, ErrInvalid
				}

				sum := 0.0
				for i := 0; i < x.row; i++ {
					sum += x.At(i) * y.At(i)
				}
				xs = append([]interface{}{sum}, xs[2:]...)
			case Mat:
				y, _ := xs[1].(Mat)
				if y.col != 1 {
					return nil, ErrInvalid
				}

				m := NewMat(x.row, y.row)
				for i := 0; i < x.row; i++ {
					for j := 0; j < y.row; j++ {
						sum := 0.0
						for k := 0; k < y.col; k++ {
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
				for i := 0; i < pro.col; i++ {
					for j := 0; j < pro.row; j++ {
						tmp := pro.At(i, j) * y
						pro.Write(i, j, tmp)
					}
				}
				xs = append([]interface{}{pro}, xs[2:]...)

			case Vec:
				y, _ := xs[1].(Vec)

				if x.row != y.row {
					return nil, ErrInvalid
				}
				pro := NewVec(x.col)
				for i := 0; i < x.col; i++ {
					sum := 0.0
					for k := 0; k < y.row; k++ {
						sum += x.At(i, k) * y.At(k)
					}
					pro.Write(i, sum)
				}
				xs = append([]interface{}{pro}, xs[2:]...)
			case Mat:
				y, _ := xs[1].(Mat)

				if x.row != y.col {
					return nil, ErrInvalid
				}
				pro := NewMat(x.col, y.row)

				for i := 0; i < x.col; i++ {
					for j := 0; j < y.row; j++ {
						sum := 0.0
						for k := 0; k < x.row; k++ {
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

func (v *Vec) Average() float64 {
	tmp := 0.0
	for i := 0; i < v.row; i++ {
		tmp += v.At(i)
	}

	return tmp / float64(v.row)
}

func (m *Mat) Average() float64 {
	tmp := 0.0
	for i := 0; i < m.col; i++ {
		tmp += m.m[i].Average()
	}

	return tmp / float64(m.col)
}

func (v *Vec) Abs() float64 {
	tmp := 0.0
	for i := 0; i < v.row; i++ {
		tmp += math.Pow(v.At(i), 2)
	}

	return math.Sqrt(tmp)
}
