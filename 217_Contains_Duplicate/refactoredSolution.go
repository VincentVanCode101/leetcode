package main

func containsDuplicate(nums []int) bool {
	solutionMap := make(map[int]int)

	for _, num := range nums {
		if _, isPresent := solutionMap[num]; isPresent {
			return true
		}
		solutionMap[num] = 1
	}

	return false
}
