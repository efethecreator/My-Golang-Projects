package main

func atoi(s string) int {
	sonuc := 0
	isaret := 1

	for index, ch := range s {
		if index == 0 && ch == '-' {
			isaret = -1
		} else if index == 0 && ch == '+' {
			isaret = 1
		} else if ch >= '0' && ch <= '9' {
			sonuc += sonuc*10 + (int(ch) - 0)
		}
	}
	return sonuc * isaret
}
