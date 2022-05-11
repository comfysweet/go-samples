package main

import (
	"fmt"
)

func cut(arr []int, ch chan int) {
	switch len(arr) {
	case 0:
		ch <- 0
	case 1:
		ch <- arr[0]
	case 2:
		ch <- arr[0] + arr[1]
	default:
		{
			ch1 := make(chan int, 2)
			go cut(arr[:(len(arr)/2)], ch1)
			go cut(arr[(len(arr)/2):], ch1)
			a := <-ch1
			b := <-ch1
			ch <- a + b
		}
	}
}

func main() {
	ch := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	go cut(arr, ch)
	result := <-ch
	fmt.Println(result)
}
