package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
Выводит текст каждой строки, которая появляется во входных
данных более одного раза. Программа читает стандартный ввод или
список именованных файлов
*/

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(stdin *os.File, counts map[string]int) {
	input := bufio.NewScanner(stdin)
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		counts[input.Text()]++
	}
}
