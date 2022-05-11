package main

import "fmt"

func main() {
	fmt.Println(fold(func(a, b int) int { return a + b }, []int{1, 2, 3, 4}))
}

func fold(op func(a, b int) int, a []int) int {
	res := a[0]
	for _, val := range a[1:] {
		res = op(res, val)
	}
	return res
}

func op(a, b int) int {
	return a + b
}
