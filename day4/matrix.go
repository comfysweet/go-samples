package main

import (
	"fmt"
)

type Matrix struct {
	a [][]int
}

func NewMatrix(l int) *Matrix {
	m := &Matrix{}
	for i := 0; i < l; i++ {
		m.a = append(m.a, make([]int, l))
	}
	return m
}

func (m *Matrix) Fill(v int) {
	for i, row := range m.a {
		for j, _ := range row {
			m.a[i][j] = v
		}
	}
}

func (m *Matrix) Add(m2 *Matrix) *Matrix {
	m3 := NewMatrix(len(m.a))
	for i, row := range m.a {
		for j, v := range row {
			m3.a[i][j] = v + m2.a[i][j]
		}
	}
	return m3
}

func (m *Matrix) String() string {
	return fmt.Sprintf("Matrix[%v ... %v]",
		m.a[0][0],
		m.a[len(m.a)-1][len(m.a)-1])
}

func main() {
	m := NewMatrix(3)
	fmt.Println(m)
	m.Fill(4)
	fmt.Println(m)
	m2 := NewMatrix(3)
	m2.Fill(3)
	m3 := m.Add(m2)
	fmt.Println(m3)
}
