package main

func maximumSubarraySum(nums []int, k int) int64 {
	var maxSum int64 = 0
	var currentSum int64 = 0

	var left = 0

	var windowMap map[int]int = make(map[int]int)

	for right := 0; right < len(nums); right++ {

		currentElement := nums[right]
		currentSum += int64(currentElement)
		windowMap[currentElement]++

		for windowMap[nums[right]] > 1 {
			leftElement := nums[left]
			currentSum -= int64(leftElement)
			windowMap[leftElement]--
			left++

		}

		if right-left+1 == k {
			if currentSum > maxSum {
				maxSum = currentSum
			}
			leftElement := nums[left]
			currentSum -= int64(leftElement)
			windowMap[leftElement]--
			left++
		}

	}

	return maxSum

}
