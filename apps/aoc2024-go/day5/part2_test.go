package day5

import (
	"testing"
)

func TestPart2Example(t *testing.T) {
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

	result := Part2Solution(pageNumbers, rules)
	expected := 123
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2Solution(t *testing.T) {
	rules := ParsePageOrderingRules(PageOrderingRules)
	pageNumbers := ParsePageNumbers(UpdatePageNumbers)

	result := Part2Solution(pageNumbers, rules)
	expected := 6767
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
