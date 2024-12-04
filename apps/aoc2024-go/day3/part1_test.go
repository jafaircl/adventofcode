package day3

import (
	"testing"
)

func TestCollectMulOperations(t *testing.T) {
	input := "mul(10, 20) mul(30, 40)"
	expected := [][]string{{"mul(10, 20)", "10", "20"}, {"mul(30, 40)", "30", "40"}}
	result := CollectMulOperations(input)
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

func TestCollectMulOperations_Example(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	expected := [][]string{{"mul(2,4)", "2", "4"}, {"mul(5,5)", "5", "5"}, {"mul(11,8)", "11", "8"}, {"mul(8,5)", "8", "5"}}
	result := CollectMulOperations(input)
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

func TestMulFromString(t *testing.T) {
	input := "mul(10, 20) mul(30, 40)"
	expected := int64(10*20 + 30*40)
	result := MulFromString(input)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFindSolutionPart1(t *testing.T) {
	result := MulFromString(CorruptedData)
	expected := int64(165225049)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
