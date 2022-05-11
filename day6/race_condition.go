package main

import (
	"fmt"
	"sync"
)

var (
	wg1 sync.WaitGroup
	a   int
)

func fR() {
	for i := 0; i < 10_000; i++ {
		a++ // a -> r; r++; r -> a
	}
	wg1.Done()
}

func main() {
	for i := 0; i < 100; i++ {
		wg1.Add(1)
		go fR()
	}
	wg1.Wait()
	fmt.Println(a)
}
