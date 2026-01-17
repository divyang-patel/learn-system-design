package lru

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	// LeetCode Example 1
	// Input:
	// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	// Output:
	// [null, null, null, 1, null, -1, null, -1, 3, 4]

	t.Run("case 1", func(t *testing.T) {
		// 1. Initialize with capacity 2
		obj := Constructor(2)

		// 2. put(1, 1)
		obj.Put(1, 1)

		// 3. put(2, 2)
		obj.Put(2, 2)

		// 4. get(1) returns 1
		if got := obj.Get(1); got != 1 {
			t.Errorf("Get(1) = %d; want 1", got)
		}

		// 5. put(3, 3) -> evicts key 2
		obj.Put(3, 3)

		// 6. get(2) returns -1 (not found)
		if got := obj.Get(2); got != -1 {
			t.Errorf("Get(2) = %d; want -1 (evicted)", got)
		}

		// 7. put(4, 4) -> evicts key 1
		obj.Put(4, 4)

		// 8. get(1) returns -1 (not found)
		if got := obj.Get(1); got != -1 {
			t.Errorf("Get(1) = %d; want -1 (evicted)", got)
		}

		// 9. get(3) returns 3
		if got := obj.Get(3); got != 3 {
			t.Errorf("Get(3) = %d; want 3", got)
		}

		// 10. get(4) returns 4
		if got := obj.Get(4); got != 4 {
			t.Errorf("Get(4) = %d; want 4", got)
		}
	})

	t.Run("values different from keys", func(t *testing.T) {
		obj := Constructor(2)
		obj.Put(10, 100)
		obj.Put(20, 200)

		if got := obj.Get(10); got != 100 {
			t.Errorf("Get(10) = %d; want 100 (Check if you are storing Key instead of Value!)", got)
		}
	})

	t.Run("update existing key check type safety", func(t *testing.T) {
		obj := Constructor(2)
		obj.Put(1, 100)
		// Update existing key
		obj.Put(1, 200)

		// This get will panic if the update put a raw int instead of listElement struct
		if got := obj.Get(1); got != 200 {
			t.Errorf("Get(1) after update = %d; want 200", got)
		}
	})
}
