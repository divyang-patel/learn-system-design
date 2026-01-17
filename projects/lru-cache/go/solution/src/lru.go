package lru

import (
	"container/list"
	"sync"
)

// LRUCache represents size-constrained Least Recently Used cache
type LRUCache struct {
	mu  sync.Mutex // The Guard
	cap int
	l   *list.List
	m   map[int]*list.Element
}

type listElement struct {
	key   int
	value int
}

// Constructor initializes a new LRUCache with the given capacity
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   list.New(),
		m:   make(map[int]*list.Element),
	}
}

// Get looks up a key's value
func (this *LRUCache) Get(key int) int {
	this.mu.Lock()
	defer this.mu.Unlock()

	if _, ok := this.m[key]; !ok {
		return -1
	}
	// TODO: Usage update!
	// Since we accessed this item, it is now the "Most Recently Used".
	// We need to move the element to the front of the list.
	this.l.MoveToFront(this.m[key])
	return this.m[key].Value.(listElement).value
}

// Put adds a key-value pair to the cache.
// If the key already exists, update the value.
// If the cache is full, evict the least recently used item.
func (this *LRUCache) Put(key int, value int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	// 1. Check if key already exists
	//    If yes: Update the value in the list element, and MoveToFront.
	if _, ok := this.m[key]; ok {
		this.m[key].Value = listElement{key, value}
		this.l.MoveToFront(this.m[key])
		return
	}

	// 2. If valid new key, but Cache is Full (this.l.Len() == this.cap):
	//    a. Find the Least Recently Used item (this.l.Back())
	//    b. Remove it from the list.
	//    c. CRITICAL: Remove it from the map `this.m`.
	//       (To do this, the list element MUST store the Key. Does it?)
	if this.l.Len() == this.cap {
		back := this.l.Back()
		this.l.Remove(back)
		delete(this.m, back.Value.(listElement).key)
	}

	// 3. Add the new item
	//    a. PushFront( {key, value} )
	//    b. Add to map `this.m`
	this.l.PushFront(listElement{key, value})
	this.m[key] = this.l.Front()
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
