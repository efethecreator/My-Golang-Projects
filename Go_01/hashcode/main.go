package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("muraytefe"))
}

func HashCode(dec string) string {
	result := ""

	for _, ch := range dec {
		result += string((int(ch) + len(dec)) % 127)
	}
	return result
}
