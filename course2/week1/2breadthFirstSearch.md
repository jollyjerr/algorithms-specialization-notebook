# Breadth First Search

> Explore nodes in layers

Can be used to compute shortest path and connected components of undirected graphs.

```psydo
breadthFirstSearch(G graph, s StartVertex) { // All nodes are initially unexplored
    let Q = <queue data structure> initialized with S
    while Q is not empty:
        remove the first node of Q, call it v
        for each edge (v, w):
            if w unexplored:
                mark w as explored
                add w to Q // at the end 
                // (first time this condition is met will be the first thing out at the next while execution)
}
```

If this is a directed graph, you need a directed path for a valid edge in the for loop

Running time of the main while loop == O(ns, ms), where

ns == # of nodes reachable from s

ms == # of edges reachable from s

## Finding shortest paths (BFS guarantees you can find this)

With respect to starting node S - the fewest hops to V.

Extra code:

```math
initialize dist(v) = 0 if v==s otherwise +Inf

when considering edge (v,w):
    if w unexplored, then set dist(w) = dist(v) + 1
```

This also corresponds to what layer it is found on.

## Undirected connectivity

Connected components = the "pieces" of a graph `G`

### Implementation

To compute all components in an undirected case

```math
start with all nodes unexplored [labelled 1 to n]

counter = 0
for i=1 to n
    if i !explored: // This includes if explored in some previous BFS
        breadthFirstSearch(G, i)
        counter++

coutnter == number of connected components
```

running time is `O(m /*o(1) per node */ + n /*o(1) per edge in each BFS*/)`
