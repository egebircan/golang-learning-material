package arrays

func Sum(array []int) (sum int) {
	for _, value := range array {
		sum += value
	}
	return
}

// Go can let you write variadic functions that can take a variable number of arguments.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(tailsToSum ...[]int) []int {
	var tailSums []int

	for _, numbers := range tailsToSum {
		if len(numbers) == 0 {
			tailSums = append(tailSums, 0)
		} else {
			tail := numbers[1:]
			tailSums = append(tailSums, Sum(tail))
		}
	}
	return tailSums
}
