# Linked Lists (Doubly Linked) in Go

## What is it?
Imagine a group of people holding hands. 
- In an **Array**, everyone stands shoulder-to-shoulder in a rigid line. To add someone in the middle, everyone else has to shuffle over 1 step.
- In a **Linked List**, people can stand anywhere in the room (memory). Each person just holds the hand of the next person.

## Types of Linked Lists

### 1. Singly Linked List (One-Way Street)
Each node holds **one** hand: the person next to them (`Next`).
- **Structure**: `[Data | Next] -> [Data | Next] -> nil`
- **Pros**: Uses less memory (only one pointer per node).
- **Cons**: You cannot go backwards. If you are at Node B, you don't know who Node A is.
- **Analogy**: A Conga line. You only know who's in front of you.

### 2. Doubly Linked List (Two-Way Street)
Each node holds **two** hands: the person before them (`Prev`) and after them (`Next`).
- **Structure**: `nil <- [Prev | Data | Next] <-> [Prev | Data | Next] -> nil`
- **Pros**: Can traverse forward and backward. efficient deletion from middle (because you can "look back" to connect your predecessor to your successor).
- **Cons**: Uses more memory (two pointers per node).
- **Analogy**: Holding hands in a chain.

### 3. Circular Linked List
The last node connects back to the first node.
- **Use Case**: Round-robin Schedulers (Player 1 -> Player 2 -> Player 1).

---

## Go's `container/list`
Go provides a standard Double Linked List.

### Visualizing Operations

**1. Initial State**
`[A] <-> [B] <-> [C] (Tail)`

**2. Move B to Front (Access B)**
`[B] <-> [A] <-> [C]`

**3. Add D (PushFront)**
`[D] <-> [B] <-> [A] <-> [C]`

**4. Evict (Remove Back)**
`[D] <-> [B] <-> [A]` (C is gone)
