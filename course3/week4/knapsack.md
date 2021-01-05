# The Knapsack Problem

Input: n items. Each has:

- value vi (nonnegative)

- size wi (nonnegative and integral)

- capacity W (nonnegative integer)

Output: A subset of the items that maximizes value of items without exceeding the capacity W with sum of sizes

## A Dynamic Solution

Notation: Let `Vix` = value of the best solution that:

1. uses only the first i items

2. has total size <= x

`Vix = max{(V[i-1],x), (v[i] + V[i-1],x-wi)}`

Edge case if wi > x, we are forced into the first case

### Subproblems

1. All possible prefixes of items

2. All possible (integral) residual capacities (when pealing away at subproblems must also peal away at W)

```psyudo
let A = 2DArray
initialize A[0, x] = 0 for 0,1,2,...W

for i = 1,2,3,...n:
    for x = 0,1,2,3,....W:
        if wi > x:
            A[i,x] := A[i-1,x]
        else:
            A[i,x] := max{A[i-1,x] , A[i-1, x-wi] + vi}
```
