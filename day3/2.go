package main

import "fmt"

func main() {
	fmt.Println(join(",", "hello", "world"))
}

func join(del string, str ...string) string {
	var res string
	for i, val := range str {
		if i != len(str)-1 {
			res += val + del
		} else {
			res += val
		}
	}
	return res
}
