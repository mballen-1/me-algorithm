package src

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParseInputString(nums []string) (n []int64, targetNumber int64, err error) {
	numsToInt := make([]int64, 0)
	fallbackArray := make([]int64, 0)
	for i, n := range nums {
		n = replaceAllWhiteSpaces(n)
		if n == "" {
			continue
		}
		num, err := strconv.ParseInt(n, 10, 64)
		numsToInt = append(numsToInt, num)
		if err != nil {
			fmt.Printf("Error parsing input at %v position\n", i+1)
			return fallbackArray, 0, err
		}
	}
	if len(numsToInt) == 0 {
		return numsToInt, 0, errors.New("Empty number list")
	}
	return numsToInt, numsToInt[len(numsToInt)-1], err
}

func replaceAllWhiteSpaces(n string) string {
	n = strings.ReplaceAll(n, " ", "")
	n = strings.ReplaceAll(n, "\r", "")
	n = strings.ReplaceAll(n, "\t", "")
	return strings.ReplaceAll(n, "\n", "")
}
