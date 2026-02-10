package kthlargest

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-Heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	k       int
	minHeap IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := IntHeap(nums)
	heap.Init(&h) // Heapify: O(N)

	for len(h) > k {
		heap.Pop(&h)
	}

	return KthLargest{
		k:       k,
		minHeap: h,
	}
}

func (this *KthLargest) Add(val int) int {
	// Standard Logic: Push, then Pop if too big.
	// heap.Push(&this.minHeap, val)
	// if len(this.minHeap) > this.k {
	// 	heap.Pop(&this.minHeap)
	// }

	// Optimization (Your Idea):
	// If heap is full (size k) and val > minHeap[0]:
	//   this.minHeap[0] = val
	//   heap.Fix(&this.minHeap, 0)
	// This is faster than Push+Pop because it avoids one re-balance.
	if len(this.minHeap) < this.k {
		heap.Push(&this.minHeap, val)
	} else if val > this.minHeap[0] {
		this.minHeap[0] = val
		heap.Fix(&this.minHeap, 0)
	}

	return this.minHeap[0]
}
