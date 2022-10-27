package code

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int      { return len(h) }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}
func Heap_code1() {
	h := &IntHeap{2, 7, 55, 1}
	heap.Init(h)
	heap.Push(h, 4)
	fmt.Println((*h)[0])

	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
