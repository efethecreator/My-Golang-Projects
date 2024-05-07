package main

import (
	"fmt"
)

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}

func NotDecimal(dec string) string {
	result := ""

	if dec == "" {
		return "\n"
	}

	for i := 0; i <= len(dec)-1; i++ {
		if dec[i] == '0' && dec[i+1] == '.' {
			continue
		} else if dec[i] >= 'a' && dec[i] <= 'z' || (dec[i] >= 'A' && dec[i] <= 'Z') {
			return dec + "\n"
		} else if dec[i] == '.' {
			continue
		} else {
			result += string(dec[i])
		}
	}
	return result + "\n"
}
