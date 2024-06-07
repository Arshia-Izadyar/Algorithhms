package sort

func Qs(arr []int, hi, lo int) {
	if lo >= hi {
		return
	}
	pivotIdx := Partition(arr, hi, lo)

	Qs(arr, pivotIdx-1, lo)
	Qs(arr, hi, pivotIdx+1)

}

func Partition(arr []int, hi, lo int) int {
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
