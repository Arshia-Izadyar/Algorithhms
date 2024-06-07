package main

import (
	"fmt"

	"github.com/Arshia-Izadyar/Algorithhms/src/sort"
)

func bubbleSort(arr []int) []int {

	for i := range arr {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr

}

func qs(arr []int, hi, lo int) {
	if lo >= hi {
		return
	}
	pivodIdx := partition(arr, hi, lo)
	qs(arr, pivodIdx-1, lo)
	qs(arr, hi, pivodIdx+1)
}

func partition(arr []int, hi, lo int) int {
	pivot := arr[hi]
	idx := lo - 1
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

func binarySearch(arr []int, val int) int {
	lo := 0
	hi := len(arr)
	for {
		if lo >= hi {
			break
		}
		mid := (hi + lo) / 2
		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			hi = mid - 1
		} else if arr[mid] < val {
			lo = mid + 1
		}
	}
	return -1
}

func selectionSor(arr []int) []int {
	n := len(arr)
	for i := range arr {
		min := i
		for j := i + 1; j < n; j++ {
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

func insertionSort(arr []int) []int {

	for i := range arr {
		for j := range arr {
			if arr[j] > arr[i] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func shellSort(arr []int) []int {
	n := len(arr)
	gap := 1
	for {
		if gap > n/3 {
			break
		}
		gap = gap*3 + 1

	}
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

func main() {
	// res := sort.ShellSort(sort.Arr)
	res := insertionSort(sort.Arr)
	fmt.Println(res)
}

// func main() {
// 	// right 2i + 2
// 	// left 2i + 1
// 	a := []int{2, 4, 5, 6, 7, 8, 9, 10}
// 	h := heap.NewMinHeap()
// 	h.Data = a
// 	h.Len = len(a) - 1

// 	h.Insert(1)
// 	h.Insert(99)
// 	h.Insert(3)
// 	fmt.Println(h.Data)
// }

// func main() {
// 	n5 := tree.BinaryNode{
// 		Value: 5,
// 		Left:  nil,
// 		Right: nil,
// 	}
// 	n4 := tree.BinaryNode{
// 		Value: 4,
// 		Left:  nil,
// 		Right: nil,
// 	}
// 	n3 := tree.BinaryNode{
// 		Value: 1,
// 		Left:  nil,
// 		Right: nil,
// 	}
// 	n2 := tree.BinaryNode{
// 		Value: 2,
// 		Left:  &n4,
// 		Right: &n5,
// 	}
// 	n1 := tree.BinaryNode{
// 		Value: 1,
// 		Left:  &n2,
// 		Right: &n3,
// 	}

// 	nn5 := tree.BinaryNode{
// 		Value: 5,
// 		Left:  nil,
// 		Right: nil,
// 	}
// 	nn4 := tree.BinaryNode{
// 		Value: 4,
// 		Left:  nil,
// 		Right: nil,
// 	}
// 	nn3 := tree.BinaryNode{
// 		Value: 3,
// 		Left:  nil,
// 		Right: nil,
// 	}
// 	nn2 := tree.BinaryNode{
// 		Value: 2,
// 		Left:  &nn4,
// 		Right: &nn5,
// 	}
// 	nn1 := tree.BinaryNode{
// 		Value: 1,
// 		Left:  &nn2,
// 		Right: &nn3,
// 	}
// res := tree.PreOrderSearch(n1)
// fmt.Println(res)
// res = tree.InOrderSearch(n1)
// fmt.Println(res)
// res = tree.PostOrderSearch(n1)
// fmt.Println(res)
// res := tree.BFSBT(&n1, 33)
// 	res := tree.Compare(&nn1, &n1)
// 	fmt.Println(res)
// }

// func main() {

// var noneSortedList = []int{99, 3, 420, 4, 2, 33, 22, 23, 9, 5, 1}
// qs(noneSortedList, len(noneSortedList)-1, 0)
// fmt.Println(noneSortedList)
// var sortedList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 23, 34, 45}
// res := binarySearch(sortedList, 45)
// fmt.Println(sortedList[res])
// l := ds.NewDoublyLinkedList[int]()
// l.Prepend(10)
// l.Prepend(11)
// l.Prepend(12)
// l.Prepend(13)

// l.InsertAt(333, 2)
// l.Append(99)
// l.Append(999)
// len := l.Len
// for i := 0; i < len; i++ {
// 	if l.Head != nil {
// 		// fmt.Println(l.Head)
// 		fmt.Printf("[%d] %d\n", i, l.Head.Value)
// 	}
// 	l.Head = l.Head.Next
// }
// res := l.Remove(13)
// fmt.Println(*res)
// l.RemoveAt(5)
// l.RemoveAt(0)
// len := l.Len
// fmt.Println(len)

// l.InsertAt(420, 2)
// // fmt.Println(l.Head)
// len = l.Len
// fmt.Println(len)
// for i := 0; i < len; i++ {
// 	if l.Head != nil {
// 		// fmt.Println(l.Head)
// 		fmt.Printf("[%d] %d\n", i, l.Head.Value)
// 		l.Head = l.Head.Next
// 	}
// 	// fmt.Println(l.Head)
// }
// }

// func main() {

// var sortedList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 23, 34, 45}

// res := search.BinarySearch(sortedList, 23)
// fmt.Println(res)

// balls := []bool{false, false, false, false, false, false, true, true, true, true, true, true}
// index := search.Two(balls)
// fmt.Println(index)

// var noneSortedList = []int{99, 3, 4, 2, 22, 9, 5, 1}

// sort.Qs(noneSortedList, len(noneSortedList)-1, 0)
// fmt.Println(noneSortedList)
// li := sort.BubbleSort(noneSortedList)
// fmt.Println(li)

// n3 := ds.Node[int]{
// 	Value: 3,
// 	Next:  nil,
// }
// n2 := ds.Node[int]{
// 	Value: 2,
// 	Next:  &n3,
// }
// n1 := ds.Node[int]{
// 	Value: 1,
// 	Next:  &n2,
// }

// ds.Remove(3, &n1)
// fmt.Println(n1.Next)
// n1.Append(12)
// n1.Append(13)
// n1.Append(13)
// fmt.Println(n1.Length())
// n1.Append(13)
// n1.Append(16)

// for {
// 	if n1.Next == nil {
// 		fmt.Println(n1.Value)
// 		break
// 	}
// 	fmt.Println(n1.Value)
// 	n1 = *n1.Next
// }

// node4 := stack.Node[int]{
// 	Value: 4,
// 	Prev:  nil,
// }

// node3 := stack.Node[int]{
// 	Value: 3,
// 	Prev:  &node4,
// }

// node2 := stack.Node[int]{
// 	Value: 2,
// 	Prev:  &node3,
// }

// node1 := stack.Node[int]{
// 	Value: 1,
// 	Prev:  &node2,
// }

// s := stack.NewStack[int]()
// s.Push(12)
// s.Push(13)
// s.Push(14)
// len := s.Length

// for i := 0; i < len; i++ {
// 	fmt.Println(s.Pop())
// 	// fmt.Println()

// }
// fmt.Println(s.Peak())
// q := queue.Queue[int]{
// 	Length: 4,
// 	Head:   &node1,
// 	Tail:   &node4,
// }
// q2 := queue.Queue[int]{

// 	Head: nil,
// 	Tail: nil,
// }
// a := q.Len()
// _ = a

// fmt.Println(*q.Peak())
// fmt.Println(q2.Peak())
// poped := q.Dequeue()
// fmt.Println(*poped)

// q.Enqueue(12)

// q.Len()
// _ = q2
//}
