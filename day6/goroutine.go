package main

import (
	"fmt"
)

func f() {
	fmt.Print("hello, ")

}

func main() {
	go f()
	fmt.Println("world!")
}
