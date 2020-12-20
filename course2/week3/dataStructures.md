# Data Structures

When learning data structures learn

1. What operations does it support?

2. What running time can you expect from those operations?

To ensure your algorithms run as fast as possible, pick a "minimal" data structure. One that only has as many features as you need.

As a general rule - the more features a data structure has, the slower it will be

## Heap (Priority queue)

### Supports two operations

1. Insertion

2. Extract min (or max)

Container for objects with a key -> use the key to compare objects

### Running time

O(logN) (With good constants)

### Other operations

Heapify O(N) batch insert

Delete (remove arbitrary)

> heaps come in handy when app is doing repeated minimum calculations and searching for them

#### Min heap property

Parent key is smaller than children key (for min heap)

## Binary Search Tree

> Like a sorted array that supports insertions and deletions < O(N^2)

### Operations

search O(logn)

select O(logn)

min/max O(logn) (search for -infinity or +infinity)

pred/succ O(logn)

rank O(logn) (augment node to store subtree size)

output in sorted order O(n) (in order traversal)

insert O(logn)

delete O(logn) (search and delete - move predecessor or single child into place)

### Structure

Each node in tree has key + pointer to some specific information

Exactly one node per key. Each node points to left child, right child, and parent

#### Search tree property

All nodes stored to the left of a node have smaller key, all nodes to the right of a node have bigger key

Guarantees you can search from root on down and forget about half the tree each time

Note: height could be anywhere from logn to n. logn is for perfectly balanced and n-1 is worst case of a chain

Note: for implementations, have a convention for collisions on insert

## Balanced search tree (Red black tree)

### Invariants

1. Store if node is "red" or "black" on each node

2. Root is always "black"

3. Never allow two reds in a row

4. Every root=>null path passes through exactly the same amount of black nodes

