package search

// func main() {

// 	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 7, 5, 4, 5, 3, 2}
// 	res := linearSearch(arr, 4)
// 	fmt.Println(res)
// 	res = linearSearch(arr, 5)
// 	fmt.Println(res)
// 	res = linearSearch(arr, 99)
// 	fmt.Println(res)
// }

func LinearSearch(arr []int, value int) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}
