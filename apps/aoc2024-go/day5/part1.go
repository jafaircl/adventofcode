package day5

func IndexOf(haystack []int64, needle int64) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func RowMedianIndex(row []int64) int {
	return len(row) / 2
}

func RowMedianValue(row []int64) int64 {
	return row[RowMedianIndex(row)]
}

func RowConformsToRule(row []int64, rule []int64) bool {
	colIndex1 := IndexOf(row, rule[0])
	// If the first page is not in the row, skip the rule
	if colIndex1 == -1 {
		return true
	}
	colIndex2 := IndexOf(row, rule[1])
	// If the second page is not in the row, skip the rule
	if colIndex2 == -1 {
		return true
	}
	return colIndex1 <= colIndex2
}

func RowConformsToPageOrderingRules(row []int64, rules [][]int64) bool {
	for _, rule := range rules {
		if !RowConformsToRule(row, rule) {
			return false
		}
	}
	return true
}

func Part1Solution(input, rules [][]int64) int {
	result := 0
	for _, row := range input {
		if RowConformsToPageOrderingRules(row, rules) {
			result += int(RowMedianValue(row))
		}
	}
	return result
}
