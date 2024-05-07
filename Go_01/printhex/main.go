package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	n := atoi(os.Args[1])

	hex := "0123456789abcdef"

	var kalan int

	var result string

	for n > 0 {
		kalan = n % 16
		n = n / 16
		result = string(hex[kalan]) + result
	}
	printstr(result)
}

func atoi(str string) int {
	sonuc := 0

	for _, ch := range str {
		if ch >= '0' && ch <= '9' {
			sonuc = sonuc*10 + int(ch-'0')
		}
	}
	return sonuc
}

func printstr(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
