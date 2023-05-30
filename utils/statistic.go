package utils

// TODO 1: Import sync package

type StatisticsResult struct {
	Mean   float64
	Median float64
	Max    int
}

// Statistics calculates thM mean, median and max of a slice of integers.
func Statistics(data []int) StatisticsResult {
	// TODO 2: Create Variables wg with type sync.WaitGroup

	// TODO 3: Create Variables mean, median with type chan float64, and max with type chan int

	// TODO 4: Create a goroutine to calculate the mean of data

	// TODO 5: Create a goroutine to calculate the median of data

	// TODO 5.1: Get Median if data is odd

	// TODO 5.2: Get Median if data is even

	// TODO 6: Create a goroutine to calculate the max of data

	// TODO 7: Create a goroutine to close all channels

	// TODO 8: Create a variable result with type StatisticsResult

	// TODO 9: Assign the value of mean, median, max to result

	// TODO 10: Return result
	return StatisticsResult{}
}
