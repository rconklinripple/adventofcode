package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(addUpValuesFromInputFile())
}

func getFirstAndLastAsTwoDigit(input string) int {
	firstNum, lastNum := firstAndLast(input)
	return firstNum*10 + lastNum
}

func firstAndLast(input string) (first int, last int) {
	for pos := 0; pos < len(input); pos++ {
		if input[pos] >= '0' && input[pos] <= '9' {
			val, err := strconv.Atoi(string(input[pos]))
			if err != nil {
				panic(err)
			}
			if first == 0 {
				first = val
			}
			last = val
		}
		wordValue := getNumberWord(input, pos)
		if wordValue >= 0 {
			if first == 0 {
				first = wordValue
			}
			last = wordValue
		}
	}
	return
}

func getNumberWord(input string, pos int) int {
	//get the current substring from our current position
	restOfString := input[pos:]
	switch {
	case strings.Index(restOfString, "one") == 0:
		return 1
	case strings.Index(restOfString, "two") == 0:
		return 2
	case strings.Index(restOfString, "three") == 0:
		return 3
	case strings.Index(restOfString, "four") == 0:
		return 4
	case strings.Index(restOfString, "five") == 0:
		return 5
	case strings.Index(restOfString, "six") == 0:
		return 6
	case strings.Index(restOfString, "seven") == 0:
		return 7
	case strings.Index(restOfString, "eight") == 0:
		return 8
	case strings.Index(restOfString, "nine") == 0:
		return 9
	case strings.Index(restOfString, "zero") == 0:
		return 0
	case strings.Index(restOfString, "two") == 0:
		return 2
	}
	return -1
}

func addUpValuesFromInputFile() (total int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total += getFirstAndLastAsTwoDigit(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}
