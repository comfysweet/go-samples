package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		log.Fatal("There are not right count of arguments")
	}

	a := parse(os.Args[1])
	b := parse(os.Args[2])
	c := parse(os.Args[3])

	if !isAllRight(a, b, c) {
		log.Fatal("Values of sides are incorrect")
	}
	p := (a + b + c) / 2
	s := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	fmt.Printf("Result: %f\n", s)
}

func isAllRight(a, b, c float64) bool {
	if a+b > c && b+c > a && c+a > b {
		return true
	}
	return false
}

func parse(s string) (res float64) {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal("Error while converting")
		return
	}
	return
}
