package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if value, ok := m[target-num]; ok {
			return []int{value, i}
		}

		m[num] = i
	}

	return nil
}
