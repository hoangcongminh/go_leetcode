package main

func search(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mid := lo + ((hi - lo) / 2)
		pivot := nums[mid]

		if pivot == target {
			return mid
		} else if pivot > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return -1
}
