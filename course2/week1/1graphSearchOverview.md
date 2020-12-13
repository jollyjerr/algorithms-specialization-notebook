# Graph Search

Graph search is an abstract process to go from an initial state to some desired state. Like finding the shortest drive or filling in a sudoku puzzle.

We will mainly cove Breadth First Search and Depth First Search.

Graph search algorithms run insanely fast, frequently giving you lots of "for free primitives".

If you have a graph and need to make sense of it, **Run these linear O(m+n) for free primitives!**.

## Graph Search Goals

1. Find everything findable from a given start index

2. Don't explore anything twice (so that you can achieve O(m+n) linear time)

## Generic Graph Search

Given graph `G` and vertex `S`

```math
Initially S is explored, all other nodes are unexplored (store a bool 'explored' with each vertex)

while possible:
    chose an edge (u,v) with u explored and v unexplored
    mark v explored
```

Claim: In an undirected or directed graph, at the end of this algorithm the nodes that have been explored are precisely the ones that can be reached via a path from `S`

The difference between graph search algorithms is how they choose what 'frontier' nodes to suck in next

## Implementations

BFS -> Explores new nodes in "layers" `O(m+n) using a que (First in, First out)` (Connected components of undirected graph)

DFS -> Explores as deep as possible. Only backtracking when necessary. `O(m+n) using a stack (Last in, First out)` (Connected components of directed graph)
