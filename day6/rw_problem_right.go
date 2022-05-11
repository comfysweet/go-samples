package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	excl        sync.Mutex
	readers     sync.Mutex
	readerCount int
)

func writer(i int) {
	for {
		excl.Lock()
		fmt.Println(i, "starts writing")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
		fmt.Println(i, "ends writing")
		excl.Unlock()
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Microsecond)
	}
}

func reader(i int) {
	for {
		readers.Lock()
		if readerCount == 0 {
			excl.Lock()
		}
		readerCount++
		readers.Unlock()

		fmt.Println(i, "starts reading")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
		fmt.Println(i, "ends reading")

		readers.Lock()
		readerCount--
		if readerCount == 0 {
			excl.Unlock()
		}
		readers.Unlock()

		time.Sleep(time.Duration(rand.Intn(3000)) * time.Microsecond)
	}
}

func main() {
	for i := 0; i < 5; i++ {
		go reader(i)
		go writer(i)
	}
	for {
	}
}
