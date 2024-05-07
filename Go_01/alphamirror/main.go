package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
	}

	sonuc := ""

	for _, char := range os.Args[1] {
		if char >= 'a' && char <= 'z' {
			opposite := 'z' - (char - 'a')
			sonuc += string(opposite)
		} else if char >= 'A' && char <= 'Z' {
			opposite := 'Z' - (char - 'A')
			sonuc += string(opposite)
		} else {
			sonuc += string(char)
		}
	}
	for _, char := range sonuc {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
