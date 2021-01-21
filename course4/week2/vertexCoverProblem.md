# Vertex Cover Problem

Input: Undirected graph G=(V,E)

Goal: Compute a minimum-cardinality vertex cover - a subset S<=V that contains at least one endpoint of each edge of G

This is in fact an NP-Complete problem.

1. Identify computationally tractable special cases

- Trees [application of dynamic programming]

- bipartite graphs [application of the maximum flow problem]

- when the optimal solution is "small" (around logn or less)

2. heuristics (suitable greedy algorithms)

3. exponential time but better than brute-force search

## An exponential time solution

Suppose: given a positive integer k as input, we want to check whether or not there is a vertex cover with size <= k. [think of k as "small"]

Could try all possibilities, would take around On^k

**Can form a smarter approach with a substructure lemma:**

Consider graph G, edge (u,v) in G, integer k >= 1

Let Gu = G with u and its incident edges deleted (similarly, Gv) =>

then G has a vertex cover of size K if and only if Gu or Gv has a vertex cover of size (k-1) // all the way down

```psydo
[given undirected graph G=(V,E), integer K]
[ignoring base cases where V=1 or k < 2]

1. Pick an arbitrary edge (u,v)
2. Recursively search for a vertex cover S of size (k-1) in Gu. If found, return S u |u|
3. Recursively search for a vertex cover S of size (k-1) in Gv. If found, return S u |v|

If they both fail, then you know that this graph does not have a "small" vertex cover
```

Running time O(2^k*m) which is way better than O(n^k). Has to be exponential time unless you plan on proving NP = P
