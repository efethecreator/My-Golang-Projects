package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	a := os.Args[1]
	b := os.Args[2]

	count := 0
	k := 0

	for i := 0; i < len(a); i++ {
		for j := k; j < len(b); j++ {
			if a[i] == b[j] {
				if i != len(a)-1 {
					i++
				}
				count++
				k = j + 1
				if i == len(a)-1 {
					count++
					break
				}
			}
		}
	}
	if count == len(a) {
		for _, char := range a {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
