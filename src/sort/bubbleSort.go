package sort

func MyBubbleSort(arr []int) []int {
	for _, k := range arr {
		_ = k
		for i := 0; i < len(arr); i++ {
			if i+1 < len(arr) {
				if arr[i] > arr[i+1] {
					swap := arr[i]
					arr[i] = arr[i+1]
					arr[i+1] = swap
				}
			}
		}
	}
	return arr
}

func BubbleSort(arr []int) []int {
	for i := range arr {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				swap := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = swap
			}
		}
	}
	return arr
}
