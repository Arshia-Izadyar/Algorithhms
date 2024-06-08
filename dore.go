package main

import (
	"fmt"
)

var list = []int{3, 5, 4, 8, 999, 99, 65, 12, 21, 30, 50, 57, 70, 2}

// Selection Sort
func selectionSort(arr []int) []int {

	// 1. go through array
	// 2. find the smallest one
	// 3. replace it with first unordered element in list
	for i := range arr {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
	return arr
}

func bubbleSort(arr []int) []int {
	for i := range arr {
		for j := 0; j < len(arr) - 1 - i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j] 
				arr[j] = arr[j + 1]
				arr[j + 1] = temp
			}
		}
	}

	return arr
}

func insertionSort(arr []int) []int {
	for i := range arr {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = key
	}
	return arr
}
// quick sort
func qs(arr []int, hi, lo int) {
	// find a pivot
	// traverse the array
	// if a element is bigger than pivot it should be at the right side pivot
	// if an element is lower than pivot should be at the right side

	if hi <= lo {
		return
	}
	p := partition(arr, hi, lo)
	qs(arr, p - 1, lo)
	qs(arr, hi, lo + 1)
	
}

func partition(arr []int, hi, lo int) int {
	var pivot = arr[hi]
	var idx = lo - 1
	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			temp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = temp
		}
	} 
	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot
	return idx
}

// shell sort
func shellSort(arr []int) []int {

	n := len(arr)
	gap := 1
	for {
		if gap > n / 3 {
			break
		}
		gap = gap * 3 + 1
	}
	for  {
		if gap < 1 {
			break
		}

		for i := gap; i < n; i++ {
			for j := i; j >= gap && arr[j] < arr[j - gap]; j-=gap {
				temp := arr[j]
				arr[j] = arr[j - gap]
				arr[j - gap] = temp
			}
		}


		gap /= 3
	}
	return arr
}

func main() {
	res := bubbleSort(list)
	// qs(list, len(list) - 1, 0)
	fmt.Println(res)

}
