package main

import "fmt"

// пример конвеера

func main() {
	numbers := make(chan int)
	doubles := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	go func() {
		for v := range numbers {
			doubles <- 2 * v
		}
		close(doubles)
	}()

	for v := range doubles {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
