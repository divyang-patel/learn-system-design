# System Design Document: LRU Cache

## 1. Requirements
### Functional Requirements
- Initialize with a fixed capacity.
- `Get(key)`: Retrieve value proficiently. Update usage.
- `Put(key, value)`: Insert or update value. Update usage. Evict if full.

### Non-Functional Requirements
- **Time Complexity:** $O(1)$ for both `Get` and `Put`.
- **Space Complexity:** $O(N)$ where $N$ is the capacity.

## 2. Architecture & Data Structures

### Challenge
We need to:
1.  Look up items quickly ($O(1)$) -> suggests Hash Map.
2.  Keep track of order (Access history) -> suggests methods like Moving to Head/Tail.
3.  Delete quickly from the middle (when accessing an old item) -> suggests Linked List.

### Choosing the Right Data Structure
To achieve $O(1)$ for both `Get` and `Put`, let's eliminate standard structures:

#### 1. Arrays or Vectors?
- **Access**: $O(1)$ by index, but we need access by Key. Finding a key requires scanning: $O(N)$.
- **Eviction**: Removing the oldest item (or moving a used item to the front) requires shifting all other elements: $O(N)$.
- **Verdict**: ❌ Too slow.

#### 2. Singly Linked Lists?
- **Insertion**: $O(1)$.
- **Access**: Finding a key requires traversal: $O(N)$.
- **Verdict**: ❌ Too slow.

#### 3. Hash Map (alone)?
- **Access**: $O(1)$. Perfect.
- **Ordering**: Hash maps are unordered. We cannot track "Recency". 
- **Verdict**: ❌ Insufficient.

#### 4. Queue + Hash Map?
A seemingly good idea: use a Queue for tracking age (FIFO) and a Map for looking up values.
- **Put (New Item)**: Push to Queue tail, add to Map. $O(1)$.
- **Evict**: Pop from Queue head, remove from Map. $O(1)$.
- **Get (The Problem)**: When we access an item in the *middle* of the Queue, we must move it to the *tail* (make it "Most Recently Used").
  - In a standard Queue (implemented as Singly Linked List), removing a node from the middle requires finding its predecessor. This search is $O(N)$.
- **Verdict**: ❌ $O(N)$ for `Get`.

#### 5. The Solution: Hash Map + Doubly Linked List
We need the fast lookup of a Map combined with the instant manipulation of a Linked List.
- **The Trick**: The Map doesn't just store the *value*. It stores a **Pointer** to the Linked List Node.
- **Fast Get**:
  1. Look up Key in Map -> Get Pointer ($O(1)$).
  2. Use Pointer to jump directly to the Node in memory.
  3. **Doubly Linked Magic**: Since we have the node, and it knows its `Prev` and `Next`, we can "snip" it out of the line and stitch the neighbors together. No traversal needed ($O(1)$).
     - *Note*: This is impossible in a Singly Linked List because you can't reach `Prev` from the current node.
  4. Move Node to Front ($O(1)$).
- **Verdict**: ✅ $O(1)$ for all operations.

- **Concurrency:** (Advanced) Is this thread-safe?
  - **Issue:** Standard Maps and Linked Lists are usually NOT thread-safe. Concurrent writes will crash or corrupt state.
  - **Solution:** Use synchronization primitives (like Mutexes or Locks) to guard shared state.
  - **Trade-off:** Locking reduces throughput (blocking), but ensures correctness.
  - **Reference:** See [Thread Safety Concepts](../../concepts/thread-safety/README.md).

## 4. Test Cases
- Cache Hit / Miss
- Eviction of least recently used
- Updating an existing key (does it become most recently used?)
- Zero capacity (edge case, though problem says positive)
