package day5

import (
	"strconv"
	"strings"
)

func ParsePageOrderingRules(input string) [][]int64 {
	result := [][]int64{}
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		if row == "" {
			continue
		}
		pages := strings.Split(row, "|")
		page1, err := strconv.ParseInt(pages[0], 10, 64)
		if err != nil {
			panic(err)
		}
		page2, err := strconv.ParseInt(pages[1], 10, 64)
		if err != nil {
			panic(err)
		}
		result = append(result, []int64{page1, page2})
	}
	return result
}

func ParsePageNumbers(input string) [][]int64 {
	result := [][]int64{}
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		if row == "" {
			continue
		}
		pages := strings.Split(row, ",")
		rowArr := []int64{}
		for _, page := range pages {
			pageNum, err := strconv.ParseInt(page, 10, 64)
			if err != nil {
				panic(err)
			}
			rowArr = append(rowArr, pageNum)
		}
		result = append(result, rowArr)
	}
	return result
}
