package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		defer close(naturals)
		for x := 0; x < 100; x++ {
			naturals <- x
		}
	}()

	// возведение в квадрат
	go func() {
		defer close(squares)
		for x := range naturals {
			squares <- x
		}
	}()

	// вывод
	for x := range squares {
		fmt.Println(x)
	}
}
