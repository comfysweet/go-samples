package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	NB = 5
)

var l sync.Mutex

func producer(i int, sem chan int) {
	for {
		c := rand.Intn(10)
		l.Lock()
		fmt.Printf("P%v: %v\n", i, c)
		sem <- c
		l.Unlock()
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}

func consumer(i int, sem chan int) {
	for {
		c := <-sem
		fmt.Printf("C%v: %v\n", i, c)
	}
}

func main() {
	sem := make(chan int, NB)
	for i := 0; i < NB; i++ {
		go producer(i, sem)
		go consumer(i, sem)
	}
	for {

	}
}
