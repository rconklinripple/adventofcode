package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type partNumber struct {
	value           int
	hasAdjacentPart bool
	startPos        int
	endPos          int
}

// Entry point, outputs the total of all the part numbers found in
// the input file
func main() {
	lines := readLines("day3input.txt")
	parts := parsePartNumbers(lines)
	total := 0
	for _, part := range parts {
		total += part.value
	}
	fmt.Println(total)
}

// parse the part numbers out of the list of strings, returning
// a list of part numbers
func parsePartNumbers(lines []string) (parts []partNumber) {
	var currentPart *partNumber = nil
	//parse line-by-line
	for lineNum, line := range lines {
		// character-by-character, look for numerics, if found
		// either update the currently parsing partNumber to add new
		// number positions, or ending when a non-part number is found
		// updating the last character position.   Also, while parsing
		// check aroudn to see if a special glyph indicates this is a valid
		// part and should be added to the list
		for charPos, _ := range line {
			if line[charPos] >= '0' && line[charPos] <= '9' {
				if currentPart == nil {
					currentPart = new(partNumber)
					currentPart.hasAdjacentPart = false
					currentPart.startPos = charPos
				}
				// handle adding digits to number
				newVal, _ := strconv.Atoi(line[charPos : charPos+1])
				currentPart.value = currentPart.value*10 + newVal //handle place shifting

				//look around for a glyph that is not a period or number
				for checkChar := -1; checkChar <= 1; checkChar++ {
					for checkLine := -1; checkLine <= 1; checkLine++ {
						if !currentPart.hasAdjacentPart {
							currentPart.hasAdjacentPart = isSpecialGlyph(charPos+checkChar, lineNum+checkLine, lines)
						}
					}
				}
			} else {
				// set the last position as the endPos and decide if it
				// should be added to the list
				if currentPart != nil {
					currentPart.endPos = charPos - 1
					if currentPart.hasAdjacentPart {
						parts = append(parts, *currentPart)
					}
					currentPart = nil
				}
			}
			// does the same as above for end-of line checking
			if charPos == len(line)-1 {
				if currentPart != nil {
					currentPart.endPos = charPos - 1
					if currentPart.hasAdjacentPart {
						parts = append(parts, *currentPart)
					}
					currentPart = nil
				}
			}
		}
	}
	return
}

// helper function to detect the special case rules to determine if a glyph is an engine part
func isSpecialGlyph(checkPos int, checkLine int, lines []string) bool {
	switch {
	case checkPos < 0:
		return false
	case checkLine < 0:
		return false
	case checkLine > len(lines)-1:
		return false
	case checkPos > len(lines[checkLine])-1:
		return false
	}
	var char rune = rune(lines[checkLine][checkPos])
	return !unicode.IsDigit(char) && char != '.'
}

// reads in the lines of text from the specified filename
// into a string array
func readLines(filename string) (lines []string) {
	file, _ := os.Open(filename)
	defer file.Close()

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		lines = append(lines, scanner.Text())
	}
	return
}
