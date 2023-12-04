package main

import (
	"fmt"
	"strings"
)

type Card struct {
	cardName       string
	pickedNumbers  []string
	winningNumbers map[string]bool
	numWinners     int
	points         int
}

func newCard(cardLine string) *Card {
	card := new(Card)
	cardValues := strings.Split(cardLine, ": ")
	card.cardName = cardValues[0]
	picksAndWinners := strings.Split(cardValues[1], " | ")

	//put winning numbers into a map
	card.winningNumbers = make(map[string]bool, 5)
	for _, winnerVal := range strings.Split(picksAndWinners[1], " ") {
		if len(winnerVal) == 1 {
			winnerVal = " " + winnerVal
		}
		card.winningNumbers[winnerVal] = true
	}

	for _, pickedNum := range strings.Split(picksAndWinners[0], " ") {
		if pickedNum == "" {
			continue
		}
		if len(pickedNum) == 1 {
			pickedNum = " " + pickedNum
		}
		if card.winningNumbers[pickedNum] {
			if card.points == 0 {
				card.points = 1
			} else {
				card.points = card.points * 2
			}
		}
	}
	return card
}

func main() {

	grandTotal := 0
	for _, line := range ReadLines("day4input.txt") {
		lineCard := newCard(line)
		grandTotal += lineCard.points
	}
	fmt.Println((grandTotal))
}