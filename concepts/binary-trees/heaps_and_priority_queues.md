# Heaps & Priority Queues

## 1. What is a Priority Queue?
A regular Queue is **FIFO** (First In, First Out). The person waiting the longest gets served first.
A **Priority Queue** looks at *VIP Status*. The person with the highest priority gets served first, regardless of how long they've been waiting.

### Real World Analogy
- **Queue**: A line at the grocery store.
- **Priority Queue**: The Emergency Room (ER) triage. A heart attack patient (High Priority) skips ahead of the person with a broken finger (Low Priority), even if the broken finger arrived an hour earlier.

## 2. What is a Heap?
A Heap is the most common Data Structure used to *implement* a Priority Queue. It is a binary tree that satisfies the **Heap Property**:

### Types
1.  **Min-Heap**: The parent is always *smaller* than its children. The Root is the minimum element.
    -   *Efficiently finds the Smallest item.*
2.  **Max-Heap**: The parent is always *larger* than its children. The Root is the maximum element.
    -   *Efficiently finds the Largest item.*

### Complexity
-   **Find Min/Max**: $O(1)$ (It's always at the top).
-   **Insert**: $O(\log N)$ (Add to bottom, "bubble up").
-   **Delete Min/Max**: $O(\log N)$ (Remove top, move bottom to top, "sift down").

## 3. Real World System Design Applications
Heaps are crucial anywhere "Scheduling" or "Top-K" problems exist.

### A. Task Scheduling
Kubernetes or OS Schedulers use heaps to decide which job to run next based on priority, deadline, or resource usage.

### B. Top-K Items (Leaderboards)
"Who are the top 10 users with the most points?"
-   Keep a Min-Heap of size 10.
-   When a new score comes in, if it's bigger than the root (the smallest of the top 10), replace the root.
-   This is much faster than sorting millions of users ($O(N \log K)$ vs $O(N \log N)$).

### C. Bandwidth Management
Network routers use Priority Queues to ensure Voice Over IP (VoIP) packets (latency-sensitive) are sent before File Download packets (bulk data).

### D. Dijkstra’s Algorithm (Shortest Path)
Finding the fastest route on Google Maps uses a Min-Heap to explore the "closest" unexplored intersections first.

## 4. Why NOT use a Heap?
-   **Searching**: Finding an arbitrary item in a heap is $O(N)$. It is NOT optimized for lookup, only for finding the "Best" item.
-   **Ordering**: It is not fully sorted. It only guarantees parent vs child order, not left child vs right child.
