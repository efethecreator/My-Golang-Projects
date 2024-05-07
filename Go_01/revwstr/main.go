package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	word := ""
	var words []string
	str := os.Args[1]
	for _, ch := range str {
		if ch != ' ' {
			word += string(ch)
		} else {
			words = append(words, word)
			word = ""
		}
	}

	if word != "" {
		words = append(words, word)
	}
	result := ""
	for i := len(words) - 1; i >= 0; i-- {
		result += words[i]
		if i != 0 {
			result += " "
		}
	}
	printstr(result)
}

func printstr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
