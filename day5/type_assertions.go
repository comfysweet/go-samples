package main

import (
	"fmt"
	"log"
)

//позволяет распаковывать интерфейсы

func executeTA(s interface{}) {
	value, ok := s.(X)
	if !ok {
		log.Fatal("incorrect type assertion")
	}
	fmt.Println(value.do())
}
