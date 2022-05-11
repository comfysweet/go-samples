package main

import (
	"fmt"
	"sync"
)

var (
	wg2 sync.WaitGroup
	a2  int
	m   sync.Mutex
)

func fM() {
	b := 0
	for i := 0; i < 10_000; i++ {
		b++
	}
	//for i := 0; i < 10_000; i++ {
	//	m.Lock() // wait() - семафоры
	//	a2++
	//	m.Unlock() // post()
	//}
	m.Lock()
	a2 += b
	m.Unlock()
	wg2.Done()
}

func main() {
	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go fM()
	}
	wg2.Wait()
	fmt.Println(a2)
}
