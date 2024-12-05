package day4

import (
	"strings"
)

// CreateBoardFromString creates a board from a string
func CreateBoardFromString(str string) [][]rune {
	// lower case letters only
	str = strings.ToLower(str)
	lines := strings.Split(str, "\n")
	board := make([][]rune, len(lines))
	for i, line := range lines {
		board[i] = []rune(line)
	}
	return board
}
