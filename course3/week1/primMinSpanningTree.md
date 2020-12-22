# Minimum Spanning Tree Problem

> Connect points together as cheaply as possible

Input: Adjacency list representation of an undirected graph that can have negative edge costs (weight)

Output: Minimum cost tree (i.e. sum of edge costs) also the output can not have cycles and the subgraph has to be connected

Assumptions: input graph G is connected, Edge costs are distinct

## Prim's algorithm

> Very similar to Dijkstra

```psyudo
initialize X = [s] (chosen arbitrarily) verts we have looked at
T = [] edges we have crossed

while X != V:
    pick the cheapest edge among all edges that cross the frontier
    add that edge to T
    add the head to X
```

Using heaps the running time can be `O(mlog(n))`

Every graph has (2^n) - 1 cuts

### Analysis

Given the assumptions, Prim's is guaranteed to output a minimum cut because it cannot get stuck and cannot produce cycles

The cut property: If an edge E crosses a single cut with the minimum weight possible then, by definition, it HAS to be included in THE (unique) minimum spanning tree of the graph

Iterative solution runs in O(mn) like Dijkstra

### Speed it up via heaps

You *could* use the heap to store edges with key = edge costs, but there is a better option

We want to use heaps to store vertices of V - X, so the heap should hand us the vertex that we are next going to add to X

The *key* for the vertex should be the cheapest edge crossing the frontier to that vertex

> After an extract min you have to recompute local keys!!

```psyudo
For each edge in the graph:
    if the edge newly crosses the frontier
        delete the head from the heap
        recompute key
        insert back into the heap
```

This code should only happen once per vertex added to X. So # of heap operations is at most O(m)
