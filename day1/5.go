package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) > 2 {
		log.Fatal("Give me a right string")
	}
	str := os.Args[1]
	if _, err := time.Parse("06.01.02", str); err != nil {
		log.Fatal("Send incorrect date")
	} else {
		fmt.Println(str)
	}
}
