package main

import "fmt"

// пример синхронного канала

const N = 100

var res int

func f(c chan int) {
	a := 0
	for i := 0; i < 100; i++ {
		a++
	}
	c <- a
}

func main() {
	c := make(chan int)
	for i := 0; i < N; i++ {
		go f(c)
	}
	for i := 0; i < N; i++ {
		res += <-c
	}
	fmt.Println(res)
}
