package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

		// no zero indexing
		lower, _ := strconv.Atoi(lowerAndUpper[0]) 
		upper, _ := strconv.Atoi(lowerAndUpper[1])
                lower -= 1
                upper -= 1

		letter := unparsedRule[1]

		if string(password[lower]) == letter && string(password[upper]) != letter {
			numValid += 1
		}
		if string(password[lower]) != letter && string(password[upper]) == letter {
			numValid += 1
		}

	}
	fmt.Println(numValid)
}
