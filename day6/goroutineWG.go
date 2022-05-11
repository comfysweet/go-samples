package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func fWG() {
	fmt.Print("hello, ")
	wg.Done()
}

func main() {
	wg.Add(1)
	go fWG()
	wg.Wait()
	fmt.Println("world!")
}
