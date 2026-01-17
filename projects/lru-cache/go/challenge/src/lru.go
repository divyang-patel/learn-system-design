package lru

// LRUCache represents size-constrained Least Recently Used cache
type LRUCache struct {
	// TODO: Define your data structures here
}

// Constructor initializes a new LRUCache with the given capacity
func Constructor(capacity int) LRUCache {
	return LRUCache{
		// TODO: Initialize fields
	}
}

// Get looks up a key's value
func (this *LRUCache) Get(key int) int {
	return -1
}

// Put adds a key-value pair to the cache. 
// If the key already exists, update the value.
// If the cache is full, evict the least recently used item.
func (this *LRUCache) Put(key int, value int) {

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
