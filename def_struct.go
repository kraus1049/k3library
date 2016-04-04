package k3library

import (
	"fmt"
)

type Vec struct {
	V   []float64
	Row int
}

type Mat struct {
	M        []Vec
	Col, Row int
}

type FNCVec struct {
	F   []func(float64, Vec) (float64, error)
	Row int
}

func NewVec(row int) Vec {
	if row < 0 {
		row = 0
	}

	v := make([]float64, row)
	return Vec{v, row}
}

func NewVecSet(xs ...float64) Vec {
	v := NewVec(len(xs))
	v.Set(xs)
	return v
}

func NewMat(col, row int) Mat {
	if col <= 0 {
		col = 1
	}

	m := make([]Vec, col)
	for i := range m {
		m[i] = NewVec(row)
	}
	return Mat{m, col, row}
}

func NewMatSet(xss [][]float64) Mat {
	row := 0
	for i := range xss {
		if len(xss[i]) > row {
			row = len(xss[i])
		}
	}

	m := NewMat(len(xss), row)
	m.Set(xss)
	return m
}

func NewFNCVec(row int) FNCVec {
	if row < 0 {
		row = 0
	}

	fs := make([]func(float64, Vec) (float64, error), row)
	return FNCVec{fs, row}
}

func NewFNCVecSet(fs ...func(float64, Vec) (float64, error)) FNCVec {
	fncv := NewFNCVec(len(fs))
	fncv.Set(fs...)
	return fncv
}

func (v Vec) String() string {
	return fmt.Sprint(v.V)
}

func (m Mat) String() string {
	str := ""
	for i, v := range m.M {
		if i == 0 {
			str += "\n|"
		} else {
			str += "|"
		}
		for _, item := range v.V {
			str += fmt.Sprintf(" %v", item)
		}
		if i != m.Col-1 {
			str += " |\n"
		} else {
			str += " |"
		}
	}

	return str
}

func (v *Vec) Set(x []float64) {
	if len(x) < v.Row {
		zero := make([]float64, v.Row-len(x))
		x = append(x, zero...)
	}

	v.V = x[0:v.Row]
}

func (m *Mat) Set(x [][]float64) {

	if len(x) >= m.Col {
		for i := 0; i < m.Col; i++ {
			m.M[i].Set(x[i])
		}
	} else {
		for i := 0; i < len(x); i++ {
			m.M[i].Set(x[i])
		}
	}

}

func (f *FNCVec) Set(fs ...func(float64, Vec) (float64, error)) {

	if len(fs) < f.Row {
		z := func(float64, Vec) (float64, error) {
			return 0, nil
		}
		zero := make([]func(float64, Vec) (float64, error), f.Row-len(fs))
		for i := range zero {
			zero[i] = z
		}

		fs = append(fs, zero...)
	}

	f.F = fs[0:f.Row]
}

func (v *Vec) At(i int) float64 {
	return v.V[i]
}

func (m *Mat) At(i, j int) float64 {
	return m.M[i].V[j]
}

func (f *FNCVec) At(i int) func(float64, Vec) (float64, error) {
	return f.F[i]
}

func (v *Vec) Write(i int, num float64) {
	if i >= 0 {
		v.V[i] = num
	} else {
		for j := 0; j < v.Row; j++ {
			v.V[j] = num
		}
	}
}

func (m *Mat) Write(i, j int, num float64) {
	if i >= 0 {
		if j >= 0 {
			m.M[i].V[j] = num
		} else {
			m.M[i].Write(-1, num)
		}
	} else {
		if j >= 0 {
			for k := 0; k < m.Col; k++ {
				m.M[k].V[j] = num
			}
		} else {
			idx := 0
			if m.Col < m.Row {
				idx = m.Col
			} else {
				idx = m.Row
			}

			for k := 0; k < idx; k++ {
				m.M[k].V[k] = num
			}
		}
	}
}

func (v *Vec) Copy() Vec {
	vec := NewVec(v.Row)
	for i := 0; i < v.Row; i++ {
		vec.Write(i, v.At(i))
	}
	return vec
}

func (m *Mat) Copy() Mat {
	mat := NewMat(m.Col, m.Row)
	for i := 0; i < m.Col; i++ {
		mat.M[i] = m.M[i].Copy()
	}
	return mat
}

func (f *FNCVec) Calc(t float64, x Vec) (Vec, error) {
	ans := make([]float64, f.Row)
	for i := range ans {
		tmp, err := f.F[i](t, x)
		if err != nil {
			return NewVec(0), err
		}

		ans[i] = tmp
	}

	return NewVecSet(ans...), nil
}
