package src

import (
	"fmt"
)

// FindSum2 is the algorithm to find the pairs from nums that sum up to a target number */
func FindSum2(nums []int64, target int64) map[int64]int64 {
	answers := make(map[int64]int64)
	if len(nums) <= 1 {
		fmt.Println("Array length <= 1, no pairs found")
		return answers
	}
	slicedNums := nums[0 : len(nums)-1]
	trackingMap, allNegativeNumbers := getTrackingMap(slicedNums)
	if allNegativeNumbers && target > 0 { // No solution
		return answers
	}
	// Check if complement of each number is inside map (O(n)), if true we have found a pair
	for _, n := range slicedNums {
		complement := target - n
		if _, found := trackingMap[complement]; found { // pair found
			updateAnswersMap(n, complement, &answers)
		}
	}
	return answers
}

// getTrackingMap returns a map that contains nums elements as key
func getTrackingMap(nums []int64) (map[int64]bool, bool) {
	trackingMap := make(map[int64]bool)
	allNegativeNumbers := true
	// Add each of the array numbers into the map
	for _, n := range nums {
		if _, ok := trackingMap[n]; !ok {
			trackingMap[n] = true
		}
		if n > 0 && allNegativeNumbers == true {
			allNegativeNumbers = false
		}
	}
	return trackingMap, allNegativeNumbers
}

// updateAnswersMap checks if a found pair is already included into the answers map, adds it if not
func updateAnswersMap(n, complement int64, answers *map[int64]int64) {
	ans := *answers
	if _, complementFound := ans[complement]; complementFound {
		return // avoids duplicating the pair
	} else {
		if n != complement { // nums should be different (E.g n = 0, complement =0 , target = 0 should be ignored)
			ans[n] = complement
		}
	}
}
