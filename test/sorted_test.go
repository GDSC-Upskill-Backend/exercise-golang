package test

import (
	"exercise_golang/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCasesSorted struct {
	Data           []int
	ExpectedResult []int
}

func TestSorted(t *testing.T) {
	// Table Driven Test
	testCases := []TestCasesSorted{
		{
			Data:           []int{1, 2, 3, 4, 5},
			ExpectedResult: []int{1, 2, 3, 4, 5},
		},
		{
			Data:           []int{3, 9, 5, 7, 1, 8, 2, 6, 4},
			ExpectedResult: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			Data:           []int{1, 7, 3, 9, 5, 7, 1, 8, 2, 6, 4, 9, 3, 1},
			ExpectedResult: []int{1, 1, 1, 2, 3, 3, 4, 5, 6, 7, 7, 8, 9, 9},
		},
	}

	for _, testCase := range testCases {
		// Sort the data
		t.Run("Test Sort", func(t *testing.T) {
			fmt.Println("=====================================")
			fmt.Println("Data: ", testCase.Data)
			fmt.Println()
			fmt.Println("Expected Sort: ", testCase.ExpectedResult)

			// Compare the result
			result := utils.Sorted(testCase.Data)
			fmt.Println("Sorted: ", result)
			assert.Equal(t, testCase.ExpectedResult, result, "The sorted data is not correct")

			fmt.Println()
			fmt.Println("=====================================")
		})
	}
}
