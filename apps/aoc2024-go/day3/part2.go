package day3

import (
	"regexp"
	"strconv"
)

func CollectMulOperationsWithDoDont(input string) [][]string {
	re := regexp.MustCompile(`mul\(\s?(\d{1,3}),\s?(\d{1,3})\)|do\(\)|don\'t\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	return matches
}

func MulFromStringWithDoDont(input string) int64 {
	shouldMul := true
	matches := CollectMulOperationsWithDoDont(input)
	result := int64(0)
	for _, match := range matches {
		if match[0] == "do()" {
			shouldMul = true
			continue
		}
		if match[0] == "don't()" {
			shouldMul = false
			continue
		}
		if shouldMul {
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
	}
	return result
}
