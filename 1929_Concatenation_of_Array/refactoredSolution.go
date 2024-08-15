package main

func getConcatenation(nums []int) []int {

	lengthOfNums := len(nums)
	result := make([]int, 2*lengthOfNums)

	for i := 0; i < lengthOfNums; i++ {
		result[i], result[i+lengthOfNums] = nums[i], nums[i]
	}

	return result

}
