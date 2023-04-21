package src

import (
	"reflect"
	"testing"
)

const (
	dummyTargetValue     = 22
	errorMessageTemplate = "\n Expected FindSum2(%v, 12) \n to be => %v, \n but got => %v"
)

func TestEmptyList(t *testing.T) {
	inputArray := make([]int64, 0)
	expected := make(map[int64]int64)
	result := FindSum2(inputArray, dummyTargetValue)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			errorMessageTemplate,
			inputArray,
			expected,
			result)
	}
}

func TestListWithOneElement(t *testing.T) {
	inputArray := make([]int64, 1)
	expected := make(map[int64]int64)
	result := FindSum2(inputArray, dummyTargetValue)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			errorMessageTemplate,
			inputArray,
			expected,
			result)
	}
}

func TestDefaultCase(t *testing.T) {
	inputArray := []int64{1, 9, 5, 0, 20, -4, 12, 16, 7, 12}
	expected := map[int64]int64{
		-4: 16,
		5:  7,
		0:  12,
	}
	result := FindSum2(inputArray, 12)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			errorMessageTemplate,
			inputArray,
			expected,
			result)
	}
}

func TestZeroAsTargetAndInsideNumsCase(t *testing.T) {
	inputArray := []int64{-100, 0, 100, 0}
	expected := map[int64]int64{
		-100: 100,
	}
	result := FindSum2(inputArray, 0)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			errorMessageTemplate,
			inputArray,
			expected,
			result)
	}
}

func TestJustNegativeNumbersAndPositiveTargetCase(t *testing.T) {
	inputArray := []int64{-100, -200, -400, 100}
	expected := make(map[int64]int64)
	result := FindSum2(inputArray, 100)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			errorMessageTemplate,
			inputArray,
			expected,
			result)
	}
}

func TestNoSolutionCase(t *testing.T) {
	inputArray := []int64{1, 9, 5, 0, 20, -4, 12, 16, 7, dummyTargetValue}
	expected := make(map[int64]int64)
	result := FindSum2(inputArray, dummyTargetValue)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			errorMessageTemplate,
			inputArray,
			expected,
			result)
	}
}
