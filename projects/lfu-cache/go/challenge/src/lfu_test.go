package lfu

import (
	"testing"
)

func TestLFUCache(t *testing.T) {
	// LeetCode Example 1
	// Input
	// ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
	// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
	// Output
	// [null, null, null, 1, null, -1, 3, null, -1, 3, 4]

	t.Run("leetcode example 1", func(t *testing.T) {
		lfu := Constructor(2)

		lfu.Put(1, 1) // cache=[1,1], cnt(1)=1
		lfu.Put(2, 2) // cache=[2,1], cnt(2)=1, cnt(1)=1

		if got := lfu.Get(1); got != 1 { // return 1. cache=[1,2], cnt(1)=2, cnt(2)=1
			t.Errorf("Get(1) = %d; want 1", got)
		}

		lfu.Put(3, 3) // 2 is the LFU key because cnt(2)=1 is min, cnt(1)=2. evict 2. cache=[3,1], cnt(3)=1, cnt(1)=2

		if got := lfu.Get(2); got != -1 {
			t.Errorf("Get(2) = %d; want -1 (evicted)", got)
		}

		if got := lfu.Get(3); got != 3 { // return 3. cache=[3,1], cnt(3)=2, cnt(1)=2
			t.Errorf("Get(3) = %d; want 3", got)
		}

		lfu.Put(4, 4) // Both 1 and 3 have the same freq, but 1 is LRU. evict 1. cache=[4,3], cnt(4)=1, cnt(3)=2

		if got := lfu.Get(1); got != -1 {
			t.Errorf("Get(1) = %d; want -1 (evicted)", got)
		}

		if got := lfu.Get(3); got != 3 {
			t.Errorf("Get(3) = %d; want 3", got)
		}

		if got := lfu.Get(4); got != 4 {
			t.Errorf("Get(4) = %d; want 4", got)
		}
	})
}
