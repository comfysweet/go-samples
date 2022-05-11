package main

import "fmt"

type Address struct {
	city string
}

func MakeAddress(city string) Address {
	return Address{city: city}
}

func NewAddress(city string) *Address {
	return &Address{
		city: city,
	}
}

func (a Address) getCityV() string {
	return a.city
}

func (a *Address) getCityP() string {
	return a.city
}

func (a Address) setCityV(city string) {
	a.city = city
}

func (a *Address) setCityP(city string) {
	a.city = city
}

func main() {
	ad1 := MakeAddress("London")
	ad2 := NewAddress("London")
	fmt.Println(ad1 == *ad2)
	fmt.Println(ad1.getCityV())
	fmt.Println(ad1.getCityP())
	ad1.setCityV("Kyiv")
	fmt.Println(ad1.city)
	ad2.setCityP("Kyiv")
	fmt.Println(ad2.city)

	fmt.Println((&Address{city: "London"}).getCityP())
}

//так нельзя
//type PInt *int
//func (p PInt) test(){}
