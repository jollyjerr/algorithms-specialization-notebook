# Optimal Binary Search Tree

Everything to left of a node is less (by key) and everything to the right is larger

A red-black tree is a solid general solution, but what if we know how frequently each node is to be searched for?

## Problem

Input: frequencies p1,p2,...pn for items 1,2,...,n

Goal: Compute a valid search tree that minimizes the weighted (average) search time

`C(T) = sum(itemsi)*pi*[search time for i in T]`

## Optimal substructure

We can prove that, if we choose the right root, each sub tree at every level will be optimal.

## Implementation

Note: Items in a subproblem are either a prefix *or* a suffix from the original problem

It is crucial that we solve smallest subproblems to largest

```psyudo
let A := 2D Array
for s=0 to n-1: // s controls the size of subproblems, representing j - i
    for i=1 to n: // so i + j == s
        A[i, i+s] = min{
            brute force search all posibilities (r := i) from i to i + s:
                Pk + A[i, r-1] + A[r+1,i+s] // interpret these lookups as 0 if not possible
        }
return A[1, n]
```

Draw this out with i and j being an axis to help it make sense

This runs at O(n^3) - possible to optimize to O(n^2) by piggy-backing off of smaller subproblem solutions but challenging so look it up if you need it
