package main

import "fmt"

type Some interface {
	do() int
}

type X int

//func (x X) do() int {
func (x *X) do() int {
	return int(*x)
}

func execute(s Some) {
	fmt.Println(s.do())
}

func main() {
	x := X(5)

	//когда (x X), &x и x будут считаться объектами Some
	//когда (x *X), тогда будет работать только для &x
	/*
		cannot use x (type X) as type Some in argument to execute:
			X does not implement Some (do method has pointer receiver)
	*/
	execute(&x)
	//execute(x)
	//для executeTA(&x) 2021/09/24 11:47:57 incorrect type assertion
	executeTA(x)
}
