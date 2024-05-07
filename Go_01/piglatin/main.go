package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
	}

	str := os.Args[1]

	vowels := "aeuioAEUIO"

	i := -1

	for index, ch := range str {
		for _, char := range vowels {
			if char == ch {
				i = index
				break
			}
		}
	}
	result := ""

	result = str[i:] + str[:i] + "ay\n"

	printstr(result)

	if i == -1 {
		printstr("novowels\n")
	}
}

func printstr(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
}
