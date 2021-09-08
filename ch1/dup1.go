package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
Выводит текст каждой строки, которая появляется в стандартом вводе
более одного раза, а также количество ее появлений.
*/
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n\n", n, line)
		}
	}
}
