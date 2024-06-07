package search

import (
	"math"
)

func TwoBalls(arr []bool) int {
	hi := len(arr)
	lo := 0
	for {
		if hi < lo {
			break
		}
		mid := (hi + lo) / 2
		if arr[mid] {
			for i := mid; i >= 0; i-- {
				if !arr[i] {
					i++
					return i
				}
			}

		} else {
			lo = mid + 1
		}
	}
	return -1
}

func Two(arr []bool) int {
	jumpSteps := math.Floor(math.Sqrt(float64(len(arr))))
	i := int(jumpSteps)

	for ; i < len(arr); i += int(jumpSteps) {
		if arr[i] {
			break
		}
	}

	i -= int(jumpSteps)
	for j := 0; j < int(jumpSteps) && i < len(arr); j++ {
		i++
		if arr[i] {
			return i
		}
	}
	return -1
}
