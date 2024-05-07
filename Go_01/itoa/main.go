package main

func itoa(n int) string {
	sonuc := ""
	isaret := ""

	if n < 0 {
		n = -n
		isaret = "-"
	}
  
	for n >= 0 {
		digit := n % 10
		n = n/10 
		sonuc = string(digit + '0') + sonuc
	}
	return isaret + sonuc
}