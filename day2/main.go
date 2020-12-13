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

const file_name = "puzzle_input.txt"

func main() {
	file, err := os.Open(file_name)

	if err != nil {
		log.Fatalf("failed to open %s", file_name)
	}
	defer file.Close()

	var (
		scanner  = bufio.NewScanner(file)
		numValid = 0
	)

	// constructing a map didn't work b/c there are
	// duplicate passwords grr
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ":")
		password := strings.TrimSpace(split[1])

		unparsedRule := strings.Split(split[0], " ")
		lowerAndUpper := strings.Split(unparsedRule[0], "-")

		lower, _ := strconv.Atoi(lowerAndUpper[0])
		upper, _ := strconv.Atoi(lowerAndUpper[1])
		letter := unparsedRule[1]

		re := regexp.MustCompile(letter)
		num := len(re.FindAllString(password, -1))

		if num >= lower && num <= upper {
			numValid += 1
		}

	}
	fmt.Println(numValid)

}
