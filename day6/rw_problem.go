package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	resource, result int
	mRW              sync.Mutex
	//wgRW             sync.WaitGroup
	flagReader bool
	flagWriter bool
	rFlag      bool
)

func write() {
	if !flagReader {
		mRW.Lock()
		flagWriter = true
		resource++
		flagWriter = false
		mRW.Unlock()
	}
	//wgRW.Done()
}

func read() {
	if !flagWriter && rFlag {
		mRW.Lock()
		flagReader = true
		result = resource
		flagReader = false
		mRW.Unlock()
	}
	//wgRW.Done()
}

func main() {
	go func() {
		for i := 0; i < 100000; i++ {
			rFlag = true
			//wgRW.Add(1)
			read()
		}
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			//wgRW.Add(1)
			write()
		}
	}()

	time.Sleep(time.Second * 10)
	//wgRW.Wait()
	fmt.Println(result)
}
