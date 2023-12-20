package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total = 0
	for scanner.Scan() {

		fmt.Println(scanner.Text())

		runes := []rune(scanner.Text())
		i, y := 0, len(runes)-1

		var firstDigit *int
		var lastDigit *int
		for firstDigit == nil || lastDigit == nil {

			if firstDigit == nil && unicode.IsDigit(runes[i]) {
				var f_d, _ = strconv.Atoi(string(runes[i]))
				firstDigit = &f_d
			}
			if lastDigit == nil && unicode.IsDigit(runes[y]) {
				var l_d, _ = strconv.Atoi(string(runes[y]))
				lastDigit = &l_d
			}

			i++
			y--
		}

		total = total + ((*firstDigit * 10) + *lastDigit)

	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
