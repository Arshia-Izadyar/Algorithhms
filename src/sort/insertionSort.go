package sort

func InsertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := range arr {
			if arr[i] < arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}
