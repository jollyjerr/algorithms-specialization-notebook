# Dynamic Knapsack Heuristic

Goal: For a user-specified parameter D > 0 guarantee a (1-E)-approximation.

Catch: Running time will increase as E decreases

## The approach: Rounding item values

Exactly solve a slightly incorrect, but easier, knapsack instance

If the wi's and W are integers, can solve in O(nw) time

If vi's are integers, can solve knapsack via dynamic programming in `O(n^2vmax)`

## Implementation

```math
Step 1:

round each vi down to the nearest multiple of "m"
[where m depends on E, exact value to be determined later]

Divide the results by m to get V'i's (integer)

Step 2:

Use dynamic programming to solve the knapsack instance with values v'... ,sizes w, capacity W
```

The higher `m` is - the faster algorithm you get but the worse accuracy.

### Second dynamic programming algorithm for knapsack

Assume values vi are integers

For `i=0,1,2,...,n and x=0,1,2,...,n*vmax`

define `Si,x` = minimum total size needed to achieve value >= x while using only the first i items (or `+Inf` if impossible)

```psyudo
Let A = 2DArray

base cases: A[0,x] = {
    0 if x=0,
    +Inf otherwise
}

for i = 1,2,3,...,n:
  for x = 0,1,2,...,n*vmax:
    A[i,x]= min{
        A[i-1.x],
        wi + A[i-1, x-vi]
    }

return the largest X suth that A[n,x] <= w
```

`m = Evmax/n`
