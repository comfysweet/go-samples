package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	for x := 1; x < 11; x++ {
		fmt.Println("send")
		out <- x
		time.Sleep(time.Second)
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		fmt.Println("receive")
		out <- v * v
		time.Sleep(time.Second)
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
