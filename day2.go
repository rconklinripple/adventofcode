package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// a set of tokens, used either as a bag or as a pull from a bag
type tokenSet struct {
	red   int
	green int
	blue  int
}

// game struct to associate the set of pulls to an id
type game struct {
	id    int
	pulls []tokenSet
}

// entry point, reads the bag, the games and then
// tests each one to ensure whether they are valid or not
// prints the total sum of the game ids that are valid
// at the end
func main2() {
	//games := readGames("day2testgamefile.txt")
	games := readGames("day2gamefile.txt")
	var total int
	for _, testgame := range games {
		gameBag := buildBag(testgame)
		powerset := gameBag.red * gameBag.blue * gameBag.green
		fmt.Println(testgame.id, ": ", powerset)
		total += powerset
	}
	fmt.Println(total)
}

// build a new minimal bag based on the game passed in.
// this bag will contain the minimum number of tokens to
// run a particular game
func buildBag(testgame game) (gameBag tokenSet) {
	for _, pull := range testgame.pulls {
		if pull.red > gameBag.red {
			gameBag.red = pull.red
		}
		if pull.green > gameBag.green {
			gameBag.green = pull.green
		}
		if pull.blue > gameBag.blue {
			gameBag.blue = pull.blue
		}
	}
	return
}

// test a game against a bag to see if it's valid
func isValid(testgame game, gameBag tokenSet) bool {
	for _, pull := range testgame.pulls {
		if pull.red > gameBag.red || pull.blue > gameBag.blue || pull.green > gameBag.green {
			return false
		}
	}
	return true
}

// handle reading the games from a file
func readGames(filename string) (games []game) {
	file, _ := os.Open(filename)
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

// handle the reading of the bags from the file system
func readBag() tokenSet {
	fileContent, _ := os.ReadFile("day2bag.txt")
	text := string(fileContent)

	return parseTokenset(text)
}

// parse the tokenset out of a string
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
