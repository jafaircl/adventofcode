package day5

import (
	"testing"
)

func TestPart1Example(t *testing.T) {
	rules := ParsePageOrderingRules(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13`)

	pageNumbers := ParsePageNumbers(`75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)

	result := Part1Solution(pageNumbers, rules)
	expected := 143
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart1Solution(t *testing.T) {
	rules := ParsePageOrderingRules(PageOrderingRules)
	pageNumbers := ParsePageNumbers(UpdatePageNumbers)

	result := Part1Solution(pageNumbers, rules)
	expected := 4462
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
