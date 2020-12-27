# Bellman Ford algorithm

Used to compute shortest path on graphs that can have negative edge weights.

It will either compute the shortest path of a graph or punt and tell you why the graph you passed it is np complete (negative cycles make this problem unsolvable)

## Under the dynamic programming paradigm

> optimal solutions must necessarily be composed of optimal solutions to smaller sub-problems

```psyudo
Let A = 2-DArray indexed by i and v. (v is the vertex and i is the limit on how many edges we can traverse for a given subproblem)

Base Case: A[0, s] = 0 || A[0, v] = +Inf for all v != s

for i := 0,1,2....n-1:
    for each v:
        A[i, v] = min{
            A[i-1, v] ||
            best of the canidates that hop to v
        }
```

### Checking for negative cycles

Run the bellman ford algorithm one extra time and if a value changes then a negative cycle must exist

### Space optimization

Delete all but the most recent round of sub problems. Store a predecessor pointer on each V so that you can reconstruct the path if needed. 

This reduces space from O(N^2) to O(N)
