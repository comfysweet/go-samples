package main

import "fmt"

// параллельно вычислить сумму элементов, используя рекурсию

func sum(a []int, c chan int) {
	s := 0
	for _, v := range a {
		s += v
	}
	c <- s
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	c1 := make(chan int)
	c2 := make(chan int)
	go sum(arr[:len(arr)/2], c1)
	go sum(arr[len(arr)/2:], c2)
	x, y := <-c1, <-c2

	fmt.Println(x + y)
}

