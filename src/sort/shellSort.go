package sort

import "fmt"

func ShellSort(arr []int) []int {
	n := len(arr)

	gap := 1
	for {
		gap = gap*3 + 1
		if gap > n/3 {
			break
		}
	}
	fmt.Println(gap)
	for {
		if gap < 1 {
			break
		}
		for i := gap; i < n; i++ {
			for j := i; j >= gap && arr[j] < arr[j-gap]; j -= gap {
				temp := arr[j]
				arr[j] = arr[j-gap]
				arr[j-gap] = temp
			}
		}
		gap /= 3
	}
	return arr
}
