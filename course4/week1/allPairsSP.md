# All Pairs Shortest Path (APSP)

## Problem

Input: Directed graph G =(V, E) with edge costs and no distinguished source vertex

Goal: Compute the shortest path for all pairs of vertices OR report a negative cycle

### Floyd-Warshall

```go-ish
let A := 3D array (indexed by i,j,k)

Pre-process base cases(
    for all i,j E v:
        A[i,j,k=0] = {
            0 if i=j,
            cij if (i,j) in e,
            +Inf if i != j and (i,j) not in e
        }
)

for k = 1 to n:
    for i = 1 to n:
        for j=1 to n:
            A[i,j,k] = min{
                A[i,j,k-1], // inherit from last step
                A[i,k,k-1] + A[k,j,k-1] // adopt this node
            }
```

O(N^3) running time

#### To check for negative cost cycle

```go-ish
A[i,i,n] for at least one value will have a negative number
```

#### Reconstruction

Keep track of B[i,j] = max label of an internal node on a shortest i-j path for all i,j in V.

Then, recursively find the path by putting together max labels of recursively smaller subproblems.

## Johnsons Algorithm

### Reweighting Technique

We cannot add constants to graphs with negative edges, but we can employ this reweighting technique:

`for every edge e=(u,v) of G, C'e := Ce + pu - pv`

C is the weight and C' the altered weight

## walkthrough

input: G

1. add a new vertex `s` which has a direct path TO all nodes with weight 0 (new graph is G')

2. Run Bellman-Ford on G` with source vertex S

3. If bFord reports a negative cost cycle - bail and report the same

4. Use the computed shortest path distances as new VERTEX weights - pv

5. Redefine edge lengths with `for every edge e=(u,v) of G, C'e := Ce + pu - pv` - They will now all be >= 0 !!! :)

6. For each v, use Dijkstra's algorithm for each possibility - this will give us shortest paths with respect to the modified costs (d')

7. Convert these modified shortest paths to the original shortest paths and return with `d(u,v) := d'(u,v) - pu + pv`

Running time: O(n) + O(mn) + O(m) + O(nmlogn) + O(n^2) = O(mnlogn)!
