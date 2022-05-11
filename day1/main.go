package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("There are not enough arguments")
	}

	for _, arg := range os.Args[1:4] {
		i, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatal("Error while converting")
		}
		fmt.Printf("%v: %[1]T\n", i)
	}
	a, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		log.Fatal("Error while converting")
	}
	fmt.Printf("%v: %[1]T\n", a)
}
