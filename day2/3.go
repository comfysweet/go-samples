package main

import "fmt"

func main() {
	str := "ddffatnkff"
	count := make(map[string]int)
	for _, r := range []rune(str) {
		count[string(r)] = count[string(r)] + 1
	}
	fmt.Println(count)
}

