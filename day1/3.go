package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	if len(os.Args) > 2 {
		log.Fatal("There are not right count of arguments")
	}
	val, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Error while converting")
	}
	switch val {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		log.Fatal("error value")
	}
}
