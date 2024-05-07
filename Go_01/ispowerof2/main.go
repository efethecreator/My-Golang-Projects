package main

func main() {
	if len(os.Args) != 2 {
		return 
	} else if len(os.Args) == 3 {
		z01.PrintRune('\n')
	}
	sayi := os.Args[1]
	atoi(sayi)
	
	result := ispowerof2(sayi) 
	if result == true {
		str("true")
	} else {
		str("false")
	}

}



func str(str string){
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}



func ispowerof2(n int) bool {
	if n == 1 {
		return true
	} else if n % 2 != 0 {
		return false 
	}
	return ispowerof2(n/2)
} 



func atoi(str string) int {
	result := 0
	for _, char := range str {
		if char >= '0' && char <= '9' {
			result := 10 * result + int(char) - 48
		} 
	}  return result
	
}