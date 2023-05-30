package test

import (
	"exercise_golang/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCasesStatistic struct {
	Data           []int
	ExpectedSort   []int
	ExpectedResult utils.StatisticsResult
}

func TestMain(m *testing.M) {
	fmt.Println("Start testing...")
	m.Run()
	fmt.Println("Finish testing...")
}

func TestStatistic(t *testing.T) {
	// Table Driven Test
	testCases := []TestCasesStatistic{
		{
			Data:         []int{1, 2, 3, 4, 5},
			ExpectedSort: []int{1, 2, 3, 4, 5},
			ExpectedResult: utils.StatisticsResult{
				Mean:   3,
				Median: 3,
				Max:    5,
			},
		},
		{
			Data:         []int{3, 9, 5, 7, 1, 8, 2, 6, 4},
			ExpectedSort: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			ExpectedResult: utils.StatisticsResult{
				Mean:   5,
				Median: 5,
				Max:    9,
			},
		},
		{
			Data:         []int{1, 7, 3, 9, 5, 7, 1, 8, 2, 6, 4, 9, 3, 1},
			ExpectedSort: []int{1, 1, 1, 2, 3, 3, 4, 5, 6, 7, 7, 8, 9, 9},
			ExpectedResult: utils.StatisticsResult{
				Mean:   4.714285714285714,
				Median: 4.5,
				Max:    9,
			},
		},
	}

	for _, testCase := range testCases {
		// Sort the data
		t.Run("Test Sort", func(t *testing.T) {
			fmt.Println("=====================================")
			fmt.Println("Data: ", testCase.Data)
			fmt.Println()
			fmt.Println("Expected Sort: ", testCase.ExpectedSort)
			sortedData := utils.Sorted(testCase.Data)
			fmt.Println("Sorted: ", sortedData)
			fmt.Println("=====================================")
			assert.Equal(t, testCase.ExpectedSort, sortedData, "Sorted is not correct")

			fmt.Println()

			t.Run("Test Statistics", func(t *testing.T) {
				// Get the statistics result
				result := utils.Statistics(sortedData)
				fmt.Println("=====================================")
				fmt.Println("Result: ", result)
				fmt.Println("Expected Result: ", testCase.ExpectedResult)
				fmt.Println("=====================================")
				fmt.Println()

				// Compare the result with expected result
				assert.Equal(t, testCase.ExpectedResult, result, "Statistics is not correct")
			})
		})
	}

}
