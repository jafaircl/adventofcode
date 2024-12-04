package day3

import (
	"regexp"
	"strconv"
)

func CollectMulOperations(input string) [][]string {
	re := regexp.MustCompile(`mul\(\s?(\d{1,3})\s?,\s?(\d{1,3})\s?\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	return matches
}

func MulFromString(input string) int64 {
	matches := CollectMulOperations(input)
	result := int64(0)
	for _, match := range matches {
		lhs, err := strconv.ParseInt(match[1], 10, 64)
		if err != nil {
			panic(err)
		}
		rhs, err := strconv.ParseInt(match[2], 10, 61)
		if err != nil {
			panic(err)
		}
		result += lhs * rhs
	}
	return result
}
