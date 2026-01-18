package lfu

import (
	"container/list"
	"sync"
)

type Node struct {
	key   int
	value int
	freq  int
}

type LFUCache struct {
	cap     int
	minFreq int
	keyMap  map[int]*list.Element // key -> list.Element (holding *Node)
	freqMap map[int]*list.List    // freq -> List of *Node
	mu      sync.Mutex
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:     capacity,
		minFreq: 0,
		keyMap:  make(map[int]*list.Element),
		freqMap: make(map[int]*list.List),
	}
}

// Get looks up a key's value
func (this *LFUCache) Get(key int) int {
	this.mu.Lock()
	defer this.mu.Unlock()

	if _, ok := this.keyMap[key]; !ok {
		return -1
	}

	element := this.keyMap[key]
	node := element.Value.(*Node)
	val := node.value

	this.updateFreq(element)

	return val
}

// Put adds a key-value pair to the cache.
func (this *LFUCache) Put(key int, value int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	if this.cap == 0 {
		return
	}

	// 1. If key exists, update value and promote freq
	if element, ok := this.keyMap[key]; ok {
		node := element.Value.(*Node)
		node.value = value
		this.updateFreq(element)
		return
	}

	// 2. If new key, check capacity
	if len(this.keyMap) >= this.cap {
		// Evict LFU (from minFreq list)
		list := this.freqMap[this.minFreq]

		// The list in map[minFreq] is guaranteed to exist and have items if we tracked minFreq correctly.
		// Front or Back?
		// New items are added to Front (PushFront).
		// So Back is the oldest (LRU) among the LFU group.
		back := list.Back()

		nodeToRemove := back.Value.(*Node)
		delete(this.keyMap, nodeToRemove.key)
		list.Remove(back)
	}

	// 3. Insert new item (Freq always starts at 1)
	this.minFreq = 1 // Basic rule: new items reset minFreq to 1
	newNode := &Node{key: key, value: value, freq: 1}

	if this.freqMap[1] == nil {
		this.freqMap[1] = list.New()
	}
	newElement := this.freqMap[1].PushFront(newNode)
	this.keyMap[key] = newElement
}

// Helper to move node from freq N list to freq N+1 list
func (this *LFUCache) updateFreq(element *list.Element) {
	node := element.Value.(*Node)
	oldFreq := node.freq

	// 1. Remove from old freq list
	oldList := this.freqMap[oldFreq]
	oldList.Remove(element)

	// 2. Update minFreq if necessary
	// If the list we just emptied was the "minimum frequency" list,
	// then the new minimum must be oldFreq + 1
	if this.minFreq == oldFreq && oldList.Len() == 0 {
		this.minFreq++
	}

	// 3. Add to new freq list
	node.freq++
	newFreq := node.freq
	if this.freqMap[newFreq] == nil {
		this.freqMap[newFreq] = list.New()
	}
	newElement := this.freqMap[newFreq].PushFront(node)

	// 4. Update pointer in keyMap
	this.keyMap[node.key] = newElement
}
