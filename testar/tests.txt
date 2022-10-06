package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for i := '9'; i >= '0'; i-- {
		z01.PrintRune(i)
	}
}
func printStr(s string) {
	for _, e := range s {
		z01.PrintRune(e)
	}
}
func retardprint(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func FirstRune(s string) rune {
	return rune(s[0])
}
func strLen(s string) int {
	return len(s)
}
func lastRune(s string) rune {
	return rune(s[len(s)])
}
