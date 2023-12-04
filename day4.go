package main

import (
	"fmt"
	"strings"
)

func main() {
	grandTotal := 0
	for _, line := range ReadLines("day4input.txt") {
		cardValues := strings.Split(line, ": ")
		picksAndWinners := strings.Split(cardValues[1], " | ")
		cardPoints := 0

		//put winning numbers into a map
		winningNumbers := make(map[string]bool, 5)
		for _, winnerVal := range strings.Split(picksAndWinners[1], " ") {
			if len(winnerVal) == 1 {
				winnerVal = " " + winnerVal
			}
			winningNumbers[winnerVal] = true
		}

		for _, pickedNum := range strings.Split(picksAndWinners[0], " ") {
			if pickedNum == "" {
				continue
			}
			if len(pickedNum) == 1 {
				pickedNum = " " + pickedNum
			}
			if winningNumbers[pickedNum] {
				if cardPoints == 0 {
					cardPoints = 1
				} else {
					cardPoints = cardPoints * 2
				}
				fmt.Println((pickedNum))
			}
		}

		fmt.Println(cardPoints, line)
		grandTotal += cardPoints
	}
	fmt.Println((grandTotal))
}
