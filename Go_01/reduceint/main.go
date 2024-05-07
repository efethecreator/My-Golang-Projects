package main

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	b := a[0]
	c := a[1]

	result := Itoa(f(b, c))

	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func Itoa(n int) string {
	isaret := ""
	result := ""
	if n < 0 {
		n = -n
		isaret = "-"
	}
	for n > 0 {
		digit := n % 10
		n = n / 10
		result = string(digit+48) + result
	}
	if result == "" {
		return "0"
	}
	return isaret + result
}
