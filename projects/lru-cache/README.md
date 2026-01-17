# LRU Cache Implementation

## Problem Statement
Design a data structure that follows the constraints of a **Least Recently Used (LRU) cache**.

Implement the `LRUCache` class:
- `LRUCache(int capacity)` Initialize the LRU cache with positive size `capacity`.
- `int get(int key)` Return the value of the `key` if the key exists, otherwise return `-1`.
- `void put(int key, int value)` Update the value of the `key` if the `key` exists. Otherwise, add the `key-value` pair to the cache. If the number of keys exceeds the `capacity` from this operation, **evict** the least recently used key.

The functions `get` and `put` must each run in **O(1)** average time complexity.

### Constraints:
- 1 <= capacity <= 3000
- 0 <= key <= 104
- 0 <= value <= 105
- At most 2 * 105 calls will be made to get and put.

Reference: [LeetCode #146](https://leetcode.com/problems/lru-cache/)

### 🌟 Bonus Challenge: Thread Safety
Modify your implementation to be **Concurrent Safe**. 
- Multiple goroutines should be able to call `Get` and `Put` simultaneously without crashing.
- Reference: [Thread Safety Concepts](../../concepts/thread-safety/README.md)


## Learning Outcomes
- Data Structures: Understanding why a simple Array or Queue is insufficient.
- Composite Data Structures: Combining Hash Maps and Doubly Linked Lists.
- Pointer Manipulation in Go.
