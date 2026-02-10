# System Design Document: K(th) Largest Element

## 1. Requirements
- **Goal**: Always be able to return the $K$-th largest number seen so far.
- **Input**: A continuous stream of numbers.
- **Constraint**: Efficient `add(val)` operation.

## 2. Architecture & Data Structures

### Analysis
If we have a stream: `[4, 5, 8, 2]` and $K=3$.
Sorted: `[2, 4, 5, 8]`. The 3rd largest is `4`.

If we add `3`: `[2, 3, 4, 5, 8]`. The 3rd largest is `4`.
If we add `10`: `[2, 3, 4, 5, 8, 10]`. The 3rd largest is `5`.

### Choosing the Right Data Structure

#### 1. Sort the entire array ($O(N \log N)$)
- Every time we `add(val)`, we resort the list.
- **Verdict**: ❌ Too slow for a stream.

#### 2. Max-Heap ($O(1)$ access max)
- Gives us the *Largest* element instantly.
- To find the $K$-th, we'd have to pop $K-1$ elements, get the result, and put them back.
- **Verdict**: ❌ Clunky.

#### 3. Min-Heap of size K (The Winner)
- We only care about the "Top K" largest elements.
- Anything smaller than the "Top K" doesn't matter.
- **Structure**: Maintain a **Min-Heap** that contains exactly $K$ elements.
    - The Root of the Min-Heap is the *Smallest* of the *Top K*.
    - By definition, the smallest of the top K largest numbers **IS** the K-th largest number.
- **Operation**:
    - If specific heap size < K: Push `val`.
    - If heap size == K:
        - If `val` > Root: Pop Root (remove smallest), Push `val`.
        - If `val` <= Root: Ignore `val` (it's too small to be in Top K).
- **Complexity**: $O(\log K)$ for `add`. $O(1)$ for `get` (peek root).
- **Verdict**: ✅ Efficient.

## 3. Key Decisions
- **Heap Implementation**: Use Go's `container/heap`.
- **Initialization**: The constructor might get an array larger than K. We need to process it to keep only K largest.

## 4. Test Cases
- `k=3`, `nums=[4, 5, 8, 2]` -> Add `3` -> Return `4`.
- `k=1`, `nums=[]` -> Add `1` -> Return `1`.
- `k=2`, `nums=[0]` -> Add `-1` -> Return `-1`? No wait, array has 2 elements now `[-1, 0]`. 2nd largest is `-1`.
