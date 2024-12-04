package day3

import (
	"testing"
)

func TestCollectMulOperationsWithDoDont(t *testing.T) {
	input := "mul(10, 20) mul(30, 40) do() don't()"
	expected := [][]string{{"mul(10, 20)", "10", "20"}, {"mul(30, 40)", "30", "40"}, {"do()", "", ""}, {"don't()", "", ""}}
	result := CollectMulOperationsWithDoDont(input)
	if len(result) != len(expected) {
		t.Errorf("Expected %d matches, got %d", len(expected), len(result))
	}
	for i, match := range result {
		for j, part := range match {
			if part != expected[i][j] {
				t.Errorf("Expected %s, got %s", expected[i][j], part)
			}
		}
	}
}

func TestMulFromStringWithDoDont(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	expected := int64(48)
	result := MulFromStringWithDoDont(input)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFindSolutionPart2(t *testing.T) {
	result := MulFromStringWithDoDont(CorruptedData)
	expected := int64(108830766)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
