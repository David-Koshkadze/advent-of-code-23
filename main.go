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

		number := getNumbers(scannerTxt)

		sum += number
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(sum)
}
