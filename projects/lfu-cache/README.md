# LFU Cache Implementation

## Problem Statement
Design and implement a data structure for a **Least Frequently Used (LFU)** cache.

Implement the `LFUCache` class:
- `LFUCache(int capacity)` Initializes the object with the capacity of the data structure.
- `int get(int key)` Gets the value of the key if the key exists in the cache. Otherwise, returns `-1`.
- `void put(int key, int value)` Update the value of the key if present, or inserts the key if not already present. When the cache reaches its capacity, it should invalidate and remove the **least frequently used** key before inserting a new item. For this problem, when there is a tie (i.e., two or more keys have the same frequency), the **least recently used** key among them should be invalidated.

To determine the least frequently used key, a **use counter** is maintained for each key in the cache. The key with the smallest use counter is the least frequently used key.

When a key is first inserted into the cache, its use counter is set to 1 (due to the put operation). The use counter for a key in the cache is incremented a an item is used (i.e., `get` or `put` operations).

The functions `get` and `put` must each run in **O(1)** average time complexity.

### Constraints:
- `1 <= capacity <= 10^4`
- `0 <= key <= 10^5`
- `0 <= value <= 10^5`
- At most `2 * 10^5` calls will be made to get and put.

### 🌟 Bonus Challenge: Thread Safety
Modify your implementation to be **Concurrent Safe**. 
- Multiple goroutines should be able to call `Get` and `Put` simultaneously without crashing.
- Reference: [Thread Safety Concepts](../../concepts/thread-safety/README.md)

Reference: [LeetCode #460](https://leetcode.com/problems/lfu-cache/)

## Learning Outcomes
- Advanced Data Structures: Managing two dimensions of ordering (Frequency AND Recency).
- Complex State Management: Synchronizing Maps and Lists.
