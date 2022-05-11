package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const NA = 10

var (
	a   int64
	wgA sync.WaitGroup
)

func fA() {
	for i := 0; i < 100; i++ {
		atomic.AddInt64(&a, 1)
	}
	wgA.Done()
}

func main() {
	wgA.Add(NA)
	for i := 0; i < NA; i++ {
		go fA()
	}
	wgA.Wait()
	fmt.Println(a)
}
