package day5

func RemoveItemFromSlice(arr []int64, index int) []int64 {
	return append(arr[:index], arr[index+1:]...)
}

func InsertItemIntoSlice(arr []int64, index int, item int64) []int64 {
	return append(arr[:index], append([]int64{item}, arr[index:]...)...)
}

func Part2Solution(input, rules [][]int64) int {
	result := 0
	for _, row := range input {
		// If the row already conforms, skip it
		if RowConformsToPageOrderingRules(row, rules) {
			continue
		}
		// Reorder the row to conform to the rules
		for {
			for _, rule := range rules {
				if !RowConformsToRule(row, rule) {
					colIndex1 := IndexOf(row, rule[0])
					colIndex2 := IndexOf(row, rule[1])
					if colIndex1 > colIndex2 {
						row = InsertItemIntoSlice(row, colIndex2, row[colIndex1])
						row = RemoveItemFromSlice(row, colIndex1+1)
					}
				}
			}
			// Once the row conforms, break the loop
			if RowConformsToPageOrderingRules(row, rules) {
				break
			}
		}
		result += int(RowMedianValue(row))
	}
	return result
}
