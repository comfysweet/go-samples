package main

import (
	"fmt"
	"os"
	"strings"
)

// выводит аргументы командной строки

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
