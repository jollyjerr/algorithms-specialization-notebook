# Clustering problems with min spanning tree

Clustering: given n "points", classify into coherent groups

## Max-Spacing k-clusterings

Assume: We know `k := # of clusters desired`

Call points p and q *separated* if they are assigned to different clusters.

Definition: The spacing of a k-clustering is the min separated distance, the bigger the better

## A greedy solution

Fuse the closest pair of separated points (or collapse clusters)

```psyudo
initially, each point is in a separate cluster
until there are only k clusters:
    let p,q := closest pair of separated points
    merge the clusters containing p and q
```

Just like kruskal's algorithm, but stopped early

(points are vertices, distances are edge costs, point pairs are edges)

This is called single-link clustering
