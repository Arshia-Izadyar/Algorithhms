package sort

var Arr []int = []int{99, 3, 420, 4, 2, 33, 22, 23, 9, 5, 1}

func SelectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
			temp := arr[i]
			arr[i] = arr[min]
			arr[min] = temp
		}
	}
	return arr
}
