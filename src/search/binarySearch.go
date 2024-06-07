package search

func BinarySearch(arr []int, value int) bool {
	high := len(arr) - 1
	low := 0

	for {
		if high < low {
			break
		}

		mid := (high + low) / 2
		if arr[mid] == value {
			return true
		} else if arr[mid] > value {
			high = mid - 1
		} else if arr[mid] < value {
			low = mid + 1
		}
	}
	return false
}
