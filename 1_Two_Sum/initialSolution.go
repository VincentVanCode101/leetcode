package main

func twoSum(nums []int, target int) []int {
	alreadySeen := make(map[int]int)
	resultIndex := make([]int, 2)

	for i := 0; i <= len(nums)-1; i++ {
		rest := target - nums[i]
		index, isPresent := alreadySeen[rest]
		if isPresent {
			resultIndex[0] = index
			resultIndex[1] = i
			return resultIndex
		} else {
			alreadySeen[nums[i]] = i
		}
	}
	return resultIndex
}
