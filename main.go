package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func Hex() {
	hex_num := "42"
	num, err := strconv.ParseInt(hex_num, 16, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("decimal num: ", num)

}
