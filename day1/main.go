package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const file_name = "puzzle_input.txt"

func main() {
	expenses := readExpenseReport()
	partOne(expenses)
	partTwo(expenses)
}

func partOne(expenses []int) {
	sort.Ints(expenses)

	// find two ints that sum to 2020
	var (
		minIndex = 0
		maxIndex = len(expenses) - 1
		sum      = expenses[minIndex] + expenses[maxIndex]
	)

	for sum != 2020 && minIndex < maxIndex {
		if sum > 2020 {
			maxIndex -= 1
		} else {
			minIndex += 1
		}

		sum = expenses[minIndex] + expenses[maxIndex]
	}
	fmt.Println(expenses[minIndex], expenses[maxIndex])
	fmt.Println("product", expenses[minIndex]*expenses[maxIndex])
}

func partTwo(expenses []int) {
	expenseMap := constructExpenseMap(expenses)

	var (
		first  = 0
		second = 1
		sum    = expenses[first] + expenses[second]
	)

	for second < len(expenses) && first < second {
		difference := 2020 - sum
		if difference > 0 {
			if expenseMap[difference] == true {
				fmt.Println(expenses[first], expenses[second], difference)
				fmt.Println("product", expenses[first]*expenses[second]*difference)
				return
			}
		}
		if second < len(expenses)-1 {
			second += 1
		} else {
			first += 1
			second = first + 1
		}
		sum = expenses[first] + expenses[second]
	}
}

func constructExpenseMap(expenses []int) map[int]bool {
	expenseMap := make(map[int]bool)
	for _, expense := range expenses {
		expenseMap[expense] = true
	}
	return expenseMap
}

func readExpenseReport() []int {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("failed to open %s", file_name)
	}
	defer file.Close()

	var (
		scanner  = bufio.NewScanner(file)
		expenses []int
	)

	for scanner.Scan() {
		expense := scanner.Text()
		int_expense, _ := strconv.Atoi(expense)
		expenses = append(expenses, int_expense)
	}

	return expenses
}
