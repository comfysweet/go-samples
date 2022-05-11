package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Defer")
		if r := recover(); r != nil {
			fmt.Println("Recovered")
		}
	}()
	panic("Error")
}
