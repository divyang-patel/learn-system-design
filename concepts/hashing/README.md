# Hashing & Hash Maps

## 1. What is Hashing?
Hashing is the process of transforming any given key (like a string "User123" or a file) into a fixed-size value (usually an integer), called a **Hash**.
Think of it like a "Digital Fingerprint".

`HashFunction("hello")` -> `18392`
`HashFunction("world")` -> `48201`

**Core Property**: The output is deterministic. Passing "hello" will *always* return `18392`.

## 2. What is a Hash Map?
A Hash Map (or Hash Table) is a data structure that implements an **Associative Array**, a structure that can map keys to values.
It uses a **Hash Function** to compute an index into an array of buckets or slots, from which the desired value can be found.

### The Magic of O(1)
In a simple array, if you want to find "User123", you have to scan every element ($O(N)$).
In a Hash Map:
1.  Computers calculate `Index = Hash("User123") % ArraySize`.
2.  It jumps directly to that index in memory.
3.  **Result**: Instant lookup, regardless of how much data is stored.

## 3. Collisions
What if `Hash("Apple")` and `Hash("Banana")` both result in index `5`? This is a **Collision**.
Common resolution strategies:
1.  **Chaining**: Each bucket holds a Linked List. If two items land in bucket 5, they share it in a list.
2.  **Open Addressing**: If bucket 5 is full, find the next available bucket.

## 4. Relevance to System Design
Hashing is everywhere in distributed systems:
1.  **Caching**: Redis/Memcached are essentially giant Hash Maps.
2.  **Load Balancing**: How do you decide which server handles a request? `Hash(RequestID) % ServerCount`.
3.  **Sharding**: Distributing a massive database across many machines. `Hash(UserID) % ShardCount`.
4.  **Consistent Hashing**: A specialized form of hashing used to minimize reorganization when servers grow/shrink (crucial for DynamoDB, Cassandra).
