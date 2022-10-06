package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	// convert to []string
	s := strings.Fields(string(content))

	fmt.Println("original string: ", s)

	fmt.Println(Parser(s))
	// Output: [a b]

}

func Parser(s []string) []string {

	var results []string

	for i := len(s) - 1; i >= 0; i-- {

		switch s[i] {
		case "(hex)":
			i--
			results = append(results, Hex(s[i]))
		case "bin":
			i--
			results = append(results, Bin(s[i]))
		default:
			results = append(results, s[i])
		}
	}
	return results
}

func Hex(s string) string {
	s = strings.ReplaceAll(s, "0x", "")
	s = strings.ReplaceAll(s, "0X", "")
	content, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(int(content))
}

func ToUpper(s string) {
	for s := 'u'; s <= 'p'; s++ {
		fmt.Printf(string(s & '_'))
	}
}

func Bin(s string) string {
	content, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println(err)
		return strconv.Itoa(int(content))
	}
	fmt.Printf("Output %d", content)
}
