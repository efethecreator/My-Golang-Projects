package main

import (
	"fmt"
)

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}

func Max(a []int) int {
	max := a[1]
	for _, char := range a {
		if char > max {
			max = char
		}
	}
	return max
}
