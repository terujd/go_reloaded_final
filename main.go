package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {

		input := args[0]
		//output args[1]
		// read input files
		content, err := os.ReadFile(input)
		if err != nil {
			log.Fatal(err)
		}

		a := Parser(content)
		b := separate(a)
		c := finalizeParser(b)

		//write string to output file args[1]
		os.WriteFile(args[1], []byte(c), 0644)
	} else {
		log.Fatal("Enter input and output file")
	}
}

// functions for the main function ---------------------------------------------------

func Parser(s []byte) []string {

	// converts byte to string array
	listOfWords := strings.Fields(string(s))

	// create a variable to contain our final modified text
	var results []string
	// this is the start of the parser, loop through the text which is to be edited starting from the back
	for i := len(listOfWords) - 1; i >= 0; i-- {
		// we pick out a word from the list using []
		switch {

		case strings.Contains(listOfWords[i], "(hex)"):
			i--
			results = append(results, Hex(listOfWords[i]))

		case strings.Contains(listOfWords[i], "(bin)"):
			i--
			results = append(results, Bin(listOfWords[i]))

		case strings.Contains(listOfWords[i], "(up)"):
			i--
			results = append(results, strings.ToUpper(listOfWords[i]))

		case strings.Contains(listOfWords[i], "(low)"):
			i--
			results = append(results, strings.ToLower(listOfWords[i]))

		case strings.Contains(listOfWords[i], "(cap)"):
			i--
			results = append(results, strings.Title(listOfWords[i]))

		// now we want to check for the cases where theres a number and a modifier example "(low, 3)"
		// because we are moving backwards in the list we need to check for the modifier first
		// if the current word contains a ")" example "3)" then check what the word before it is example "(low,"
		case strings.HasSuffix(listOfWords[i], ")") && strings.Contains(listOfWords[i-1], "(low,"):

			// because the modifier number is written in this format, example "3)" we need to remove the ")" from the word

			// remove the ")" from the word and convert the number before it to a integer
			lowNum, _ := strconv.Atoi(strings.TrimSuffix(listOfWords[i], ")"))
			i--
			// converts the words before the modifier "(low," to lowercase and add them to the results
			// inside this loop we also want to move backwards in the list so we need to subtract 1 from i
			for j := 0; j < lowNum; j++ {

				i--
				results = append(results, strings.ToLower(listOfWords[i]))

				// if the list is only 5 words and if the number is 8 then we need to make sure the loop stops at the end of the list
				if i == 0 {
					break
				}
			}

		case strings.HasSuffix(listOfWords[i], ")") && strings.Contains(listOfWords[i-1], "(up,"):

			highNum, _ := strconv.Atoi(strings.TrimSuffix(listOfWords[i], ")"))
			i--
			for j := 0; j < highNum; j++ {
				i--
				results = append(results, strings.ToUpper(listOfWords[i]))
				if i == 0 {
					break
				}
			}

		case strings.HasSuffix(listOfWords[i], ")") && strings.Contains(listOfWords[i-1], "(up,"):

			highNum, _ := strconv.Atoi(strings.TrimSuffix(listOfWords[i], ")"))
			i--
			for j := 0; j < highNum; j++ {
				i--
				results = append(results, strings.ToUpper(listOfWords[i]))
				if i == 0 {
					break
				}
			}

		case strings.HasSuffix(listOfWords[i], ")") && strings.Contains(listOfWords[i-1], "(cap,"):

			capNum, _ := strconv.Atoi(strings.TrimSuffix(listOfWords[i], ")"))
			i--
			for j := 0; j < capNum; j++ {
				i--
				results = append(results, strings.Title(listOfWords[i]))
				if i == 0 {
					break
				}
			}
		default:
			results = append(results, listOfWords[i])
		}
	}
	for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
		results[i], results[j] = results[j], results[i]
	}
	return results
}

func separate(s []string) []string {

	temp1 := strings.Join(s, " ")

	for _, e := range temp1 {
		switch {
		case e == ',':
			temp1 = strings.Replace(temp1, ",", ", ", -1)
		case e == '.':
			temp1 = strings.Replace(temp1, ".", ". ", -1)
		case e == '!':
			temp1 = strings.Replace(temp1, "!", "! ", -1)
		case e == '?':
			temp1 = strings.Replace(temp1, "?", "? ", -1)
		case e == ';':
			temp1 = strings.Replace(temp1, ";", "; ", -1)
		case e == ':':
			temp1 = strings.Replace(temp1, ":", ": ", -1)
		case e == '\'':
			temp1 = strings.Replace(temp1, "'", " ' ", -1)
		}
	}
	trimWhite := regexp.MustCompile(`\s+`)
	temp1 = trimWhite.ReplaceAllString(temp1, " ")
	return strings.Split(temp1, " ")
}

func finalizeParser(results []string) string {
	var temp []string
	apostrophecount := 0
	for i := range results {

		switch {
		// if e == a and the next word starts with a vowel then add "n" to the end of the word
		case results[i] == "a" || results[i] == "A":
			if i != len(results)-1 && strings.ContainsAny(string(results[i+1][0]), "aoueihAOUEIH") {
				temp = append(temp, results[i]+"n")
			} else {
				temp = append(temp, results[i])
			}
		// if e == ",;:.?!" then add it to the end of the word before it
		case results[i] == "," || results[i] == ";" || results[i] == ":" || results[i] == "." || results[i] == "?" || results[i] == "!":
			temp[len(temp)-1] = temp[len(temp)-1] + results[i]

		case results[i] == "'":
			// if apostrophecount is 0 then add it to the beginning of the word after it
			if apostrophecount == 0 {
				results[i+1] = "'" + results[i+1]
				apostrophecount++
			} else {
				// if apostrophecount is 1 then add it to the end of the word before it
				temp[len(temp)-1] = temp[len(temp)-1] + results[i]
				apostrophecount = 0
			}
		default:
			temp = append(temp, results[i])
		}
	}
	result := strings.Join(temp, " ")
	//if theres a space in the end of the string then remove it
	if result[len(result)-1:] == " " {
		result = result[:len(result)-1]
	}
	return result

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

func Bin(s string) string {
	content, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.Itoa(int(content))
}
