package main

func containsDuplicate(nums []int) bool {

	solutionMap := make(map[int]int)

	for _, num := range nums {
		_, isPresent := solutionMap[num]
		if isPresent {
			return true
		}
		if !isPresent {
			solutionMap[num] = 1
		}
	}

	return false
}
