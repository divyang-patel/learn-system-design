package kthlargest

import (
	"testing"
)

func TestKthLargest(t *testing.T) {
	// LeetCode Example 1
	// Input
	// ["KthLargest", "add", "add", "add", "add", "add"]
	// [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
	// Output
	// [null, 4, 5, 5, 8, 8]

	t.Run("leetcode example 1", func(t *testing.T) {
		k := 3
		nums := []int{4, 5, 8, 2}
		obj := Constructor(k, nums)

		// 1. Add 3.  Array: [2, 3, 4, 5, 8], Top 3: [4, 5, 8]. Min is 4.
		if got := obj.Add(3); got != 4 {
			t.Errorf("Add(3) = %d; want 4", got)
		}

		// 2. Add 5. Array: [2, 3, 4, 5, 5, 8], Top 3: [5, 5, 8]. Min is 5.
		if got := obj.Add(5); got != 5 {
			t.Errorf("Add(5) = %d; want 5", got)
		}

		// 3. Add 10. Array: [..., 4, 5, 5, 8, 10], Top 3: [5, 8, 10]. Min is 5.
		if got := obj.Add(10); got != 5 {
			t.Errorf("Add(10) = %d; want 5", got)
		}

		// 4. Add 9. Top 3: [8, 9, 10]. Min is 8.
		if got := obj.Add(9); got != 8 {
			t.Errorf("Add(9) = %d; want 8", got)
		}

		// 5. Add 4. Top 3: [8, 9, 10] (4 is too small). Min is 8.
		if got := obj.Add(4); got != 8 {
			t.Errorf("Add(4) = %d; want 8", got)
		}
	})
}
