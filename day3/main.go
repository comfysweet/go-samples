package main

import "fmt"

func inc(x int) int {
	return x + 1
}

func double(x int) int {
	return x * 2
}

func main() {
	// такая штука не работает
	//fmt.Printf("%v", f())

	closures()

	variadic()

	//immediately-invoked
	fmt.Println(s)

	//defer
	deferExample()

}

func f() (x, y int) {
	return 1, 2
}

var s = func(x, y int) int {
	return x + y
}(2, 3)

func closures() {
	var l []func()
	for i := 0; i < 10; i++ {
		x := i
		l = append(l, func() {
			fmt.Println(x)
		})
	}
	for _, f := range l {
		f()
	}
}

func sum(vs ...int) int {
	s := 0
	for _, v := range vs {
		s += v
	}
	return s
}

func variadic() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))
	l := []int{1, 2, 3, 4}
	fmt.Println(sum(l...))
}

func deferExample(){
	defer fmt.Println("world")
	fmt.Print("Hello, ")
}