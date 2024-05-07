package main

import (
	"fmt"
)

func main() {
	fmt.Print(ReverseSecondHalf("This is the 1st half This is the 2nd half"))
	fmt.Print(ReverseSecondHalf(""))
	fmt.Print(ReverseSecondHalf("utku"))
}

func ReverseSecondHalf(str string) string {
	if str == "" {
		return "invalid\n"
	}

	if len(str) == 1 {
		return str + "\n"
	}

	sonuc := ""

	for i := len(str) - 1; i >= len(str)/2; i-- {
		sonuc += string(str[i])
	}
	return sonuc + "\n"
}
