package day4

import (
	"testing"
)

func TestFindCrossWords(t *testing.T) {
	data := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	input := CreateBoardFromString(data)
	result := FindCrosses(input)
	expected := 9
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFindPart2Solution(t *testing.T) {
	input := CreateBoardFromString(PuzzleInput)
	result := FindCrosses(input)
	expected := 2005
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
