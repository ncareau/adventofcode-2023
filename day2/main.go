package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var red_max = 12
	var green_max = 13
	var blue_max = 14

	var total = 0

	red_regexp := regexp.MustCompile("([0-9]+) red")
	green_regexp := regexp.MustCompile("([0-9]+) green")
	blue_regexp := regexp.MustCompile("([0-9]+) blue")

	for scanner.Scan() {
		isInvalid := false
		//fmt.Println(scanner.Text())

		var game_string = strings.Split(scanner.Text(), ": ")

		var game_id, _ = strconv.Atoi(strings.Split(game_string[0], " ")[1])

		var round_string = strings.Split(game_string[1], ";")

		for i := 0; i < len(round_string); i++ {

			red_matches := red_regexp.FindStringSubmatch(round_string[i])
			green_matches := green_regexp.FindStringSubmatch(round_string[i])
			blue_matches := blue_regexp.FindStringSubmatch(round_string[i])

			var round_red = 0
			if len(red_matches) > 0 {
				round_red, _ = strconv.Atoi(red_matches[1])
			}
			var round_green = 0
			if len(green_matches) > 0 {
				round_green, _ = strconv.Atoi(green_matches[1])
			}
			var round_blue = 0
			if len(blue_matches) > 0 {
				round_blue, _ = strconv.Atoi(blue_matches[1])
			}

			if round_red > red_max || round_green > green_max || round_blue > blue_max {
				isInvalid = true
			}

		}

		if isInvalid {
			continue
		}

		total = total + game_id
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
