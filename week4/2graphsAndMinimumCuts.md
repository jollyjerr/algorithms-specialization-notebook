# Graphs and minimum cuts

Graph -> represent pair wise relationships

Two ingredients:

- vertices aka nodes (V)

- edges (E) -> pairs of vertices

    - undirected (edges are unordered pair)

    - directed (ordered pair) aka arcs (from `tail` to `head`)

Cuts in graphs:

A partition of V into two non empty sets A and B

The `crossing edges` of a cut (A, B) are two with:

```math
undirected: one endpoint in each of (A, B)
directed: tail in A, head in B
```

A graph of N vertices has roughly 2^N-2 cuts

## The minimum cut problem

> Part of the genre of graph partitioning

```math
Input: an undirected graph G = (V, E)
Goal: Compute the cut with the smallest number of crossing edges. (Parallel edges are allowed)
```

### Graph partitioning

Detect weaknesses in networks. Identify communities in social networks. Image segmentation.

## Sparse an Dense Graphs

```math
let N = # of vertices, M = # of edges

In most but not all applications: m is O(N^2)
```

In a "sparse graph", m is O(N) or close to it

In a "dense graph", m is closer to O(N^2)

## Representing graphs

### The Adjacency Matrix

Represent the graph using a matrix

```math
n = number of vertices of the graph
n*n matrix A => Aij == 1 if there is an edge between i and j

Variations:
Aij number of arcs
Aij weight of ij edge
Aij +1 || -1 for directed graphs
```

Adjacency Matrix requires O(n^2) resources (with n number of vertices)

very wastefulin sparse graphs

### Adjacency Lists

```math
x = array of vertices
y = array of edges

each edge points to its endpoints
each vertex points to edges incident on it
```

Requires O(m+n) space

> Picking a data structure depends on density of graph and what kind of operations you want to perform

Adjacency List is great for graph search

