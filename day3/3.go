package main

import "fmt"

func main() {
	fmt.Println(polynome([]float64{2, 3, 4})(2.0))
}

//схема горнера
func polynome(arr []float64) func(ff float64) float64 {
	return func(x float64) float64 {
		var res float64
		for _, v := range arr {
			res = res*x + v
		}
		return res
	}
}
