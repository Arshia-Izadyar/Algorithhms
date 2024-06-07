package heap

import "fmt"

type MinHeap struct {
	Len  int
	Data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		Len:  0,
		Data: []int{},
	}
}

func (h *MinHeap) Insert(val int) {

	h.Len++
	h.Data = append(h.Data, val)
	h.hipifyUp(h.Len)

}
func (h *MinHeap) Delete(val int) int {
	if h.Len == 0 {
		return -1
	}
	var out int = h.Data[0]
	h.Len--
	if h.Len == 0 {
		h.Data = []int{}
		return out
	}

	var last int = h.Data[h.Len]
	h.Data[0] = last
	h.hipifyDown(0)
	return out

}

func (h *MinHeap) hipifyDown(idx int) {
	lIdx := h.leftChild(idx)
	rIdx := h.rightChild(idx)
	if idx >= h.Len || lIdx >= h.Len {
		return
	}
	v := h.Data[idx]
	lv := h.Data[lIdx]
	rv := h.Data[rIdx]

	if lv > rv && v > rv {
		h.Data[idx] = rv
		h.Data[rIdx] = v
		h.hipifyDown(rIdx)
	} else if lv < rv && v > lv {
		h.Data[idx] = lv
		h.Data[lIdx] = v
		h.hipifyDown(lIdx)
	}
}

func (h *MinHeap) hipifyUp(idx int) {
	if idx == 0 {
		return
	}
	fmt.Println(idx)
	var p = h.parent(idx)
	var parentValue = h.Data[p]
	var v = h.Data[idx]
	fmt.Println(parentValue)
	fmt.Println("-----------")

	if v < parentValue {
		h.Data[idx] = parentValue
		h.Data[p] = v
		h.hipifyUp(p)
	}

}

func (h *MinHeap) parent(idx int) int {
	return (idx - 1) / 2
}
func (h *MinHeap) rightChild(idx int) int {
	return idx*2 + 2
}
func (h *MinHeap) leftChild(idx int) int {
	return idx*2 + 1
}
