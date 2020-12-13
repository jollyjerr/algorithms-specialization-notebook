# Random Contraction

```psydo
while there are more than two vertices:
    (u, v) := pick a remaining edge uniformily at random()
    merge (or contract) u and v into a single vertex
    delete self loops
return cut represented by final two vertices
```

![](../images/mincutexample.png)

This algorithm can literally straight up get it wrong depending on the random choices but it is just unlikely!

## Analysis

Question: What is the probability of success?

### Setup

Fix a graph G = (v, e) with N vertices and M edges

Fix a specific minimum cut (A, B)

K = size of the minimum cut (call crossing edges F)

## What could go wrong?

If one of F edges gets chosen, the algorithm will not output (A, B)

If only edges inside A or inside B get contracted, the algorithm will not output (A, B)

`Pr[output is (A, B)] === Pr[never contracts an edge in F]`

### The first iteration

 Pr[] that we screw up on iteration one is `# of crossing edges/# of edges`

Key observation: degree (# of incident edges) of each vertex is at least K

Reason: each vertex v defines a cut ({v}, V - {v})

Sum of the degree of the vertices is exactly double edges

So Pr[] that we screw up on iteration one is `2/# of vertices`

### The second iteration

Pr[!S2 | S1] >= 1 - (2/(n-1))

### So in general

`Pr[!s1 | S2 | S3...] >= (1-2/n) * (1-2/(n-1)) * (1-2/(n-2)) *... (1-2/(n-(n-4))) * (1-2/(n-(n-3))`

And then massive cancelation happens!

`2/n(n-1)` or `1/n^2`

This leads us to see we have a low success probability - but we can boost with repeated trials

## Repeated trials

Return the best of a bunch of repeated trials

The probability becomes the product of all of your trials probability `1 - (1/n^2)^N`

`Pr[all fail] <= (e^-(1/n^2))^2` = `1/e`

Add a natural log factor = `1/n`

### Runtime

is polynomial and high but can be shaved to about O(n^2) with memoization and such

Fun fact: No graph has > N^2 number of minimum cuts
