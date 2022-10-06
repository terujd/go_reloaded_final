package main

import "github.com/01-edu/z01"

func main() {
	printStr("9876543210")

	/*for i := '9'; i >= '0'; i-- {
		z01.PrintRune(i)
	}

	z01.PrintRune('\n')
	*/
}

/*fmt.Println(strLen("hej beg"))
}

func strLen(s string) int {
	return len(s)
}*/

/*func main() {
	i := '0'
	for e := 0; e < len(os.Args[1:]); e++ {
		i++
	}
	z01.PrintRune(i)
	z01.PrintRune('\n')
}*/

func printStr(s string) {
	for _, e := range s {
		z01.PrintRune(e)
	}
}
