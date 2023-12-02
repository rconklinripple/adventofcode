package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tokenSet struct {
	red   int
	green int
	blue  int
}
type game struct {
	id    int
	pulls []tokenSet
}

func main() {
	gameBag := readBag()
	fmt.Println(gameBag)
	games := readGames()
	var total int
	for _, testgame := range games {
		if isValid(testgame, gameBag) {
			fmt.Println(testgame.id, " possible")
			total += testgame.id
		} else {
			fmt.Println(testgame.id, " impossible")
		}
	}
	fmt.Println(total)
}
func isValid(testgame game, gameBag tokenSet) bool {
	for _, pull := range testgame.pulls {
		if pull.red > gameBag.red || pull.blue > gameBag.blue || pull.green > gameBag.green {
			return false
		}
	}
	return true
}
func readGames() (games []game) {
	file, _ := os.Open("day2gamefile.txt")
	defer file.Close()

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		var readgame game
		//split into game label vs game values fields
		fields := strings.Split(scanner.Text(), ":")
		//parse out the id
		readgame.id, _ = strconv.Atoi(strings.Split(fields[0], " ")[1])

		//split into "pulls"
		pulls := strings.Split(fields[1], ";")
		for _, pullString := range pulls {
			readgame.pulls = append(readgame.pulls, parseTokenset(strings.TrimSpace(pullString)))
		}
		games = append(games, readgame)
	}
	return
}
func readBag() tokenSet {
	fileContent, _ := os.ReadFile("day2bag.txt")
	text := string(fileContent)

	return parseTokenset(text)
}

func parseTokenset(text string) (filebag tokenSet) {
	tokenTypes := strings.Split(text, ", ")
	for _, tokenType := range tokenTypes {
		fields := strings.Split(tokenType, " ")
		val, _ := strconv.Atoi(fields[0])
		switch strings.TrimSpace(fields[1]) {
		case "red":
			filebag.red = val
		case "green":
			filebag.green = val
		case "blue":
			filebag.blue = val
		}
	}
	return
}
