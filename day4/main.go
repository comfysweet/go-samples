package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

type Boss struct {
	Person
	salary float64
}

func main() {
	b := new(Boss)
	//разработчики разрешили эту * не писать
	(*b).firstName = "Harry"
	b.lastName = "Potter"
	b.salary = 99999
	fmt.Println(b)
	fmt.Println(*b)
}
