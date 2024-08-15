package main

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		sum := target - nums[i]

		if num1, ok := m[sum]; ok {
			return []int{num1, i}
		}

		m[nums[i]] = i
	}
	return []int{0, 0}
}
