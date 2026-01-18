# System Design Document: LFU Cache

## 1. Requirements
### Functional Requirements
- **Capacity**: Limit total items.
- **Eviction Policy**: 
    1. Primary: Frequency (Least Used).
    2. Secondary: Recency (Least Recently Used, if frequencies tie).
- **Get(key)**: Return value, Increment frequency.
- **Put(key, value)**: Insert/Update. Increment frequency. Evict if full.

### Non-Functional Requirements
- **Time Complexity**: $O(1)$ for both `Get` and `Put`.
- **Space Complexity**: $O(N)$ where $N$ is capacity.

## 2. Architecture & Data Structures

### Challenge
We need to track items by **ID** (for lookup) AND by **Frequency** (for eviction).
With a simple Min-Heap, `Get` would be $O(1)$ but updating the frequency in the heap would be $O(\log N)$.
We need $O(1)$.

### Choosing the Right Data Structure

#### Approach 1: Min-Heap (Priority Queue)
- **Idea**: Store all items in a Min-Heap ordered by Frequency.
- **Eviction**: `heap.Pop()` gives the LFU item. $O(\log N)$.
- **Access/Update**: When `Get(key)` happens, we must increment frequency. This changes the item's priority. We must fix the Heap.
- **Verdict**: ❌ $O(\log N)$ complexity. Fails requirements.

#### Approach 2: Linked Hash Map + Frequency Arrays (The Winner)
- **Idea**: Since frequencies are integers (1, 2, 3...), we can use them as indices (or keys in a Map).
- **Structure**:
  1. `keyMap`: `key -> Node*` (Standard fast lookup).
  2. **`freqMap`**: `frequency -> DoublyLinkedList`.
     - `freqMap[1]`: List of all items accessed once.
     - `freqMap[2]`: List of all items accessed twice.
- **Eviction**: Keep a global `minFreq` variable. Look at `freqMap[minFreq]`. Pop the tail (LRU within LFU).
- **Update**:
  1. Unlink Node from `freqMap[oldFreq]`.
  2. Link Node to `freqMap[newFreq]`.
  3. If `freqMap[minFreq]` is empty, `minFreq++`.
- **Verdict**: ✅ $O(1)$ for all operations.

## 3. Key Decisions
- **Decision 1:** ...
- **Concurrency:** Is this thread-safe?
  - **Issue:** Standard Maps and Linked Lists are usually NOT thread-safe. Concurrent writes will crash or corrupt state.
  - **Solution:** Use synchronization primitives (like Mutexes or Locks) to guard shared state.
  - **Trade-off:** Locking reduces throughput (blocking), but ensures correctness.
  - **Reference:** See [Thread Safety Concepts](../../concepts/thread-safety/README.md).

## 4. Test Cases
- **Frequency Increment**: `Put(A)`, `Put(A)` -> Freq 2.
- **Tie Breaking**: `A` (freq 1), `B` (freq 1). Cache Full. Evict A? (assuming A is older).
- **Min Freq Update**: If we upgrade the only item with freq 1 to freq 2, `minFreq` must become 2.
