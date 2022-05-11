package main

import "fmt"

func main() {
	arr := [...]int{11, 56, 33, 71, 44, 4}
	var min, max, imin, imax int

	for i, _ := range arr {
		if i == 0 {
			min = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	for i, val := range arr {
		if val == min {
			imin = i
		}
		if val == max {
			imax = i
		}
	}
	arr[imin], arr[imax] = arr[imax], arr[imin]
	fmt.Printf("new arr: %v", arr)
}
