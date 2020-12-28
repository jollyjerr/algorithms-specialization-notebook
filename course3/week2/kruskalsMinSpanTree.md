# Kruskal's minimum spanning tree

Look at the cheapest edge overall and pick it (without making cycles)

```psydo
sort edges in order of increasing cost (O(mlogn))
X=null
for i=1 to m:
    if X+i has no cycles: (bfs or dfs in O(N))
        add i to X 
return X
```

Straightforward implementation takes O(mn)

We can improve to constant time cycle checks with *Union-Find* data structure

Two operations:

1. Find(x): Return the name of the group that x belongs to

2. Union(1,2): fuse groups 1 and 2 into a single one

## Implementing union-find

> Maintain one linked structure per connected component of (V, T)

Each component has an arbitrary leader vertex

Invariant: each vertex points to the leader of it's component

If two vertices have pointers to the same leader, then they are already connected. With this implementation the find operation should return the leader name.

### In order to keep the union operation in logarithmic time

When two components merge, have smaller one inherit the leader from the larger group (maintain a size field)

O(nlogn) global upper bound for this operation
