package main

import "fmt"

func main() {
	s := "Привет, мир!"
	for _, r := range []rune(s) {
		fmt.Printf("%c\n", r)
	}
	s1 := `Multiline
		string`
	fmt.Println(s1)
}
