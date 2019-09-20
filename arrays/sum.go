package arrays

// Sum returns the sum of all the numbers inside the integer array
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

/*
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
*/

// SumAll which will take a varying number of slices, returning a new slice containing the totals for each slice passed in
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails  where it now calculates the totals of the "tails" of each slice. The tail of a collection is all the items apart from the first one (the "head")
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {

		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// Slices can be sliced! The syntax is slice[low:high]
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
