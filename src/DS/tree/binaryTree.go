package tree

import "fmt"

type BinaryNode struct {
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}

// var path []int = []int{}

func walk(curr *BinaryNode, path *[]int) {
	if curr == nil {
		return
	}

	*path = append(*path, curr.Value)
	// fmt.Println(path)
	walk(curr.Left, path)
	walk(curr.Right, path)

}

func PreOrderSearch(head BinaryNode) []int {
	p := []int{}
	walk(&head, &p)
	return p
}

func walkForInOrder(curr *BinaryNode, path *[]int) {
	if curr == nil {
		return
	}

	// fmt.Println(path)
	walkForInOrder(curr.Left, path)
	*path = append(*path, curr.Value)
	walkForInOrder(curr.Right, path)

}

func InOrderSearch(head BinaryNode) []int {
	p := []int{}
	walkForInOrder(&head, &p)
	return p
}

func walkLRV(curr *BinaryNode, path *[]int) {
	if curr == nil {
		return
	}

	// fmt.Println(path)
	walkLRV(curr.Left, path)
	walkLRV(curr.Right, path)
	*path = append(*path, curr.Value)

}

func PostOrderSearch(head BinaryNode) []int {
	p := []int{}
	walkLRV(&head, &p)
	return p
}

func BFSBT(head *BinaryNode, needle int) bool {
	q := []*BinaryNode{head}
	for {

		if len(q) == 0 {
			return false
		}

		curr := q[0]
		fmt.Println(curr.Value)
		q = q[1:]

		if curr.Value == needle {
			return true
		}
		if curr.Left != nil {

			q = append(q, curr.Left)
		}
		if curr.Right != nil {

			q = append(q, curr.Right)
		}
	}
}
