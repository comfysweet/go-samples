package main

import "fmt"

func main() {
	f(5)
	f("test")
}

func f(t interface{}) {
	switch q := t.(type) {
	case int:
		fmt.Println("int: ", q)
	case string:
		fmt.Println("string: ", q)
	default:
		fmt.Println("Unsupported type")
	}
}
