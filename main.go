package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// get numbers from the string
func getNumbers(str string) int {
	var numbers []string

	for _, char := range str {
		if unicode.IsDigit(char) {
			numbers = append(numbers, string(char))
		}
	}

	result := []string{numbers[0], numbers[len(numbers)-1]}

	strNum := strings.Join(result[:], "")

	num, err := strconv.Atoi(string(strNum))

	if err != nil {
		return -1
	}

	return num
}

// day 1 part 2

var numberWords = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getNumbersNumeric(str string) int {
	var numStr string

	words := strings.Fields(str)

	for _, word := range words {
		if num, ok := numberWords[word]; ok {
			numStr += strconv.Itoa(num)
			fmt.Println(num)
		} else {
			for _, char := range word {
				if unicode.IsDigit(char) {
					numStr += string(char)
				}
			}
		}
	}

	if len(numStr) == 0 {
		return 0
	}

	num, _ := strconv.Atoi(string(numStr[0]) + string(numStr[len(numStr)-1]))
	return num
}

func main() {
	file, err := os.Open("text_input.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int

	for scanner.Scan() {
		scannerTxt := scanner.Text()

		number := getNumbersNumeric(scannerTxt)

		// fmt.Println(number)

		sum += number
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	// fmt.Println(sum)

	fmt.Println(getNumbersNumeric("twonrpvnnmvkh2threejzcpz"))
}
