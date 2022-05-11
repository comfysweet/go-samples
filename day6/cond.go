package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// задача о producer/consumer
// решается с использованием буфера
// но буфер может либо переполниться, либо быть пустым
// нужно гарантировать, что не будет проблем с полным и пустым буфером
// нужно понять, как условные переменные помогают в решении этой проблемы

const N = 5

var (
	// циклический буффер fifo очередь
	buff              [N]int
	head, tail, count int
	l                 sync.Mutex
)

func producer(i int, notFull, notEmpty *sync.Cond) {
	for {
		l.Lock()
		for count == N {
			// должны подождать, Lock отпускаем
			// обладает свойством случайных пробуждений
			fmt.Printf("P%v: full\n", i)
			notFull.Wait()
		}
		count++
		buff[head] = rand.Intn(10)
		fmt.Printf("P%v: %v\n", i, buff[head])
		head = (head + 1) % N
		notEmpty.Signal()
		l.Unlock()
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}

func consumer(i int, notFull, notEmpty *sync.Cond) {
	for {
		l.Lock()
		for count == 0 {
			fmt.Printf("C%v: empty\n", i)
			notEmpty.Wait()
		}
		count--
		fmt.Printf("C%v: %v\n", i, buff[tail])
		tail = (tail + 1) % N
		notFull.Signal()
		l.Unlock()
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}

func main() {
	notFull, notEmpty := sync.NewCond(&l), sync.NewCond(&l)
	for i := 0; i < 5; i++ {
		go producer(i, notFull, notEmpty)
		go consumer(i, notFull, notEmpty)
	}
	for {

	}
}
