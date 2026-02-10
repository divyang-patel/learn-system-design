# Binary Trees & Hierarchies

## 1. Introduction
A **Tree** is a hierarchical data structure. Unlike Arrays or Linked Lists (which are linear), Trees are non-linear.
- **Root**: The top node.
- **Child**: A node directly connected to another node when moving away from the Root.
- **Parent**: The converse notion of a child.
- **Leaf**: A node with no children.

## 2. Binary Tree
A special type of tree where each node has **at most two children** (Left and Right).

### Types
1.  **Full Binary Tree**: Every node has 0 or 2 children.
2.  **Complete Binary Tree**: All levels are completely filled except possibly the last level, which is filled from left to right.
    -   *Crucial for Heaps!*
3.  **Binary Search Tree (BST)**:
    -   Left Child < Parent
    -   Right Child > Parent
    -   *Crucial for Searching ($O(\log N)$).*

## 3. Related Concepts
This section covers specialized tree structures used in system design:

- [Heaps & Priority Queues](./heaps_and_priority_queues.md): Specialized Complete Binary Trees used for scheduling and ordering.
- **Tries (Prefix Trees)**: For autocomplete systems.
- **B-Trees / B+ Trees**: For Database Indexing (optimizing disk reads).

## 4. Real World Application
- **DOM**: The HTML structure of a webpage is a tree.
- **FileSystem**: Folders and files form a tree.
- **Databases**: Indexes (B-Trees) allow SQL to find rows in $O(\log N)$ instead of scanning the whole table $O(N)$.
