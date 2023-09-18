package main

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	m := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}

		m[num] = struct{}{}
	}

	return false
}
