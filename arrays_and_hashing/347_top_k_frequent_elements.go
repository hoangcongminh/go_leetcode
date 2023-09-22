package main

func topKFrequent(nums []int, k int) []int {
	var res []int
	countMap := make(map[int]int)
	countList := make([][]int, len(nums)+1)

	for _, num := range nums {
		if count, ok := countMap[num]; ok {
			countMap[num] = count + 1
		} else {
			countMap[num] = 1
		}
	}

	for num, count := range countMap {
		countList[count] = append(countList[count], num)
	}

	for i := len(countList) - 1; i > 0; i-- {
		res = append(res, countList[i]...)

		if len(res) == k {
			return res
		}
	}

	return res
}
