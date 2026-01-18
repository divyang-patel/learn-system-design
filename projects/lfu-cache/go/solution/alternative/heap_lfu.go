package lfu

import (
	"container/heap"
)

// NOTE: This file uses package 'lfu' so it can participate in the same tests.
// In a real scenario, you wouldn't have two 'LFUCache' structs in the same package.
// You would need to rename this file or the struct if compiling both.

type Item struct {
	key       int
	val       int
	freq      int
	timestamp int // needed for LRU tie-breaking
	index     int
}

// PriorityQueue implements heap.Interface
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 1. Primary Sort: Frequency
	if pq[i].freq != pq[j].freq {
		return pq[i].freq < pq[j].freq
	}
	// 2. Secondary Sort: Recency (Oldest timestamp wins/evicts first)
	return pq[i].timestamp < pq[j].timestamp
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type LFUCache struct {
	cap  int
	tick int // Global logical clock
	pq   PriorityQueue
	m    map[int]*Item
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:  capacity,
		tick: 0,
		pq:   make(PriorityQueue, 0), // min-heap
		m:    make(map[int]*Item),
	}
}

func (this *LFUCache) Get(key int) int {
	if item, ok := this.m[key]; ok {
		// Update Logic
		item.freq++
		this.tick++
		item.timestamp = this.tick

		// The expensive O(log N) operation
		heap.Fix(&this.pq, item.index)
		return item.val
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}

	if item, ok := this.m[key]; ok {
		item.val = value
		item.freq++
		this.tick++
		item.timestamp = this.tick
		heap.Fix(&this.pq, item.index)
		return
	}

	// Eviction
	if len(this.pq) >= this.cap {
		// Pop the LFU item (Min-Heap root)
		item := heap.Pop(&this.pq).(*Item)
		delete(this.m, item.key)
	}

	// Insert
	this.tick++
	newItem := &Item{
		key:       key,
		val:       value,
		freq:      1,
		timestamp: this.tick,
	}
	heap.Push(&this.pq, newItem)
	this.m[key] = newItem
}
