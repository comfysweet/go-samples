package main

import (
	"fmt"
)

type Str2 struct {
	A2 int
}
type Str struct {
	A0   int
	Str1 struct {
		A1 int
	}
	Str2
}

func main() {
	var str = Str{
		A0: 1,
		Str1: struct {
			A1 int
		}{
			A1: 4,
		},
		Str2: Str2{
			A2: 3,
		},
	}
	fmt.Printf("%#v\n", str)
}
