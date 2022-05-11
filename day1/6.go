package main

import (
	"fmt"
	"log"
	"os"
)

// работает только для строк, в которых четное количество букв
func main() {
	if len(os.Args) > 2 {
		log.Fatal("Give me a right string")
	}
	str := os.Args[1]
	var begin string
	var end string
	for i := 0; i < len(str)/2; i++ {
		if str[i] == str[len(str)-i-1] {
			begin = begin + string(str[i])
			end = string(str[len(str)-i-1]) + end
		}
	}
	result := begin + end
	fmt.Printf("Current string is %s, substring is %s, Lenght is %d,", str, result, len(result))
}
