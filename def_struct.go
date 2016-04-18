package k3library

import (
	"fmt"
)

type Vec struct {
	v   []float64
	row int
}

type Mat struct {
	m        []Vec
	col, row int
}

type FNCVec struct {
	f   []func(float64, Vec) (float64, error)
	row int
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
	return fmt.Sprint(v.v)
}

func (m Mat) String() string {
	str := ""
	for i, v := range m.m {
		if i == 0 {
			str += "\n|"
		} else {
			str += "|"
		}
		for _, item := range v.v {
			str += fmt.Sprintf(" %v", item)
		}
		if i != m.col-1 {
			str += " |\n"
		} else {
			str += " |"
		}
	}

	return str
}

func (v *Vec) Set(x []float64) {
	if len(x) < v.row {
		zero := make([]float64, v.row-len(x))
		x = append(x, zero...)
	}

	v.v = x[0:v.row]
}

func (m *Mat) Set(x [][]float64) {

	if len(x) >= m.col {
		for i := 0; i < m.col; i++ {
			m.m[i].Set(x[i])
		}
	} else {
		for i := 0; i < len(x); i++ {
			m.m[i].Set(x[i])
		}
	}

}

func (f *FNCVec) Set(fs ...func(float64, Vec) (float64, error)) {

	if len(fs) < f.row {
		z := func(float64, Vec) (float64, error) {
			return 0, nil
		}
		zero := make([]func(float64, Vec) (float64, error), f.row-len(fs))
		for i := range zero {
			zero[i] = z
		}

		fs = append(fs, zero...)
	}

	f.f = fs[0:f.row]
}

func (v *Vec) V() []float64 {
	return v.v
}

func (v *Vec) Row() int {
	return v.row
}

func (m *Mat) M() []Vec {
	return m.m
}

func (m *Mat) Col() int {
	return m.col
}

func (m *Mat) Row() int {
	return m.row
}

func (f *FNCVec) F() []func(float64, Vec) (float64, error) {
	return f.f
}

func (f *FNCVec) Row() int {
	return f.row
}

func (v *Vec) At(i int) float64 {
	return v.v[i]
}

func (m *Mat) At(i, j int) float64 {
	return m.m[i].v[j]
}

func (f *FNCVec) At(i int) func(float64, Vec) (float64, error) {
	return f.f[i]
}

func (v *Vec) Write(i int, num float64) {
	if i >= 0 {
		v.v[i] = num
	} else {
		for j := 0; j < v.row; j++ {
			v.v[j] = num
		}
	}
}

func (m *Mat) Write(i, j int, num float64) {
	if i >= 0 {
		if j >= 0 {
			m.m[i].v[j] = num
		} else {
			m.m[i].Write(-1, num)
		}
	} else {
		if j >= 0 {
			for k := 0; k < m.col; k++ {
				m.m[k].v[j] = num
			}
		} else {
			idx := 0
			if m.col < m.row {
				idx = m.col
			} else {
				idx = m.row
			}

			for k := 0; k < idx; k++ {
				m.m[k].v[k] = num
			}
		}
	}
}

func (v *Vec) Copy() Vec {
	vec := NewVec(v.row)
	for i := 0; i < v.row; i++ {
		vec.Write(i, v.At(i))
	}
	return vec
}

func (m *Mat) Copy() Mat {
	mat := NewMat(m.col, m.row)
	for i := 0; i < m.col; i++ {
		mat.m[i] = m.m[i].Copy()
	}
	return mat
}

func (f *FNCVec) Calc(t float64, x Vec) (Vec, error) {
	ans := make([]float64, f.row)
	for i := range ans {
		tmp, err := f.f[i](t, x)
		if err != nil {
			return NewVec(0), err
		}

		ans[i] = tmp
	}

	return NewVecSet(ans...), nil
}
