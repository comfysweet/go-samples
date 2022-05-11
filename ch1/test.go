package main

import "fmt"

func main() {
	r := [...]int{101: -2}
	fmt.Println(r)
	fmt.Println(len(r))

	e := map[int]string{}
	fmt.Println(e)
	for _, v := range e {
		fmt.Println(v)
	}
}
