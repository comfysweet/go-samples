package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	tick := time.Tick(time.Second)
	abort := make(chan struct{})
	wg.Add(1)
	go func() {
		for {
			select {
			case <-tick:
				fmt.Println("Tick")
			case <-abort:
				fmt.Println("Abort")
				wg.Done()
				return
			}
		}
	}()
	os.Stdin.Read(make([]byte, 1))
	abort <- struct{}{}
	wg.Wait()
}
