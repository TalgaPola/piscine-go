package piscine

import (
	"github.com/01-edu/z01"
)

func IsNegative(nb int) {

	if nb<0 {
		z01.PrintRune('T')
	}
	else {
		z01.PrintRune('F')
	}
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}
