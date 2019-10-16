package piscine

import "github.com/01-edu/z01"

func get_l(a, b, c rune) rune {
	if c == a {
		return b + 1
	} else {
		return '0'
	}
}

func PrintComb2() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := a; c <= '9'; c++ {
				h := get_l(a, b, c)
				for ; h <= '9'; h++ {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(32)
					z01.PrintRune(c)
					z01.PrintRune(h)
					if a != '9' || b != '8' || c != '9' || h != '9' {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}			
				}
			}
		}
	}
	z01.PrintRune(10)
}
