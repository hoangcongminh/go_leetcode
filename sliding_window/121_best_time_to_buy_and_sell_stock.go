package main

import (
	"math"
)

func maxProfit(prices []int) int {
	res := 0
	min := math.MaxUint32

	for _, price := range prices {
		if min < price {
			if (price - min) > res {
				res = price - min
			}
		} else {
			min = price
		}

	}

	return res
}
