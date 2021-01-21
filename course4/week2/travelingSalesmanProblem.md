# Traveling Salesman Problem

Input: a complete undirected graph with nonnegative edge costs.

Output: a minimum-cost tour. A cycle that visits every vertex exactly once.

This problem is NP-Complete

Brute-Force search takes around n! time

DynamicProgramming: will obtain O(n^2*2^n) running time

## A optimal substructure lemma?

Idea: copy the format of the bellman-ford algorithm (edge budget)

For every edge budget `i` in `0,1,2,...n` with destination `j` in `1,2,...n` let `Lij` = the length of a shortest path from 1 to j that uses at most i edges

SADLY - solving all subproblems doesn't solve the original problem :(!!

Same goes for if you insist all nodes are visited once.

## Settle for a non-optimal substructure lemma

For every destination `j` in `1,2,...,n` every subset `S` <= `1,2,...,n` that contains 1 and j, let `Ls,j` = minimum length of a path from 1 to j that visits precisely the vertices of `S` (exactly once each)

Let P be a shortest path from 1 to j that visits the vertices S exactly once each. If last hop of P is (k,j), then P' is a shortest path from 1 to k that visits every vertex of S - {j} exactly once.

```psyudo
Let A = 2D array, indexed by subsets S in {1,2,...,n} that contain 1 and destination j in {1,2,...,n}

BaseCases: A[S,1] = {
    0 if S = {1},
    +Inf otherwise [no way to avoid visiting vertex 1 twice]
}

for m=2,3,4,...,n:
    for each set S in {1,2,...,n} of size m that contains 1:
        for each j in S, j != 1:
            A[S,j] = min(k in S, k!=j){
                A[S-{j}, k] + Ckj,
            }

return min(n-1 for every j){
    A[{1,2,3,...,n}, j] + Cj1
}
```
