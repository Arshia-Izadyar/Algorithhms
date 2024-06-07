package tree

// walk dfs
/*
func w(node1, node2 *BinaryNode, res *[]int) {
	if node1 == nil || node2 == nil {
		return
	}

	if node1.Value != node2.Value {
		// res = false
		*res = append(*res, node1.Value)
		*res = append(*res, node2.Value)
		return
	}
	w(node1.Left, node2.Left, res)
	w(node1.Right, node2.Right, res)

}
*/

func w(node1, node2 *BinaryNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Value != node2.Value {
		return false
	}
	return w(node1.Left, node2.Left) && w(node1.Right, node2.Right)
}

func Compare(node1, node2 *BinaryNode) bool {
	// res := []int{}
	// w(node1, node2, &res)

	res := w(node1, node2)
	return res
}
