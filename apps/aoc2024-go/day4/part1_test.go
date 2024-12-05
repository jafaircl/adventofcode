package day4

import (
	"testing"
)

func TestFindWords(t *testing.T) {
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
	result := FindWords(input, []string{"xmas"})
	expected := 18
	if result["xmas"] != expected {
		t.Errorf("Expected %d, got %d", expected, result["xmas"])
	}
}

func TestFindPart1Solution(t *testing.T) {
	input := CreateBoardFromString(PuzzleInput)
	result := FindWords(input, []string{"xmas"})
	expected := 2639
	if result["xmas"] != expected {
		t.Errorf("Expected %d, got %d", expected, result["xmas"])
	}
}
