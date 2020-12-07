# general mathematical tool for analyzing the running time of divide and conquer algorithms

## Pre-Info

The way you reason about the running time of recursive algorithms is by using a "recurrence"

```math
T(N) Worst case number of operations that the algorithm requires (the upper bound)
 
 A recurrence is to express T(N) in terms of the work done in it's recurrsive calls
```

Two ingredients:

1. Base Case `T(1 or 2 or something) >= a constant`

2. General case `T(N) <= <number of recursive calls at this step>T(N/2 <or whatever they are invoked on>) + <work done outside of recursive calls> (ex: O(N) linear)`

## Mathematical definition of the master method

Assumptions:

1. All subproblems have equal size

Recurrence format:

```math
1. BaseCase: T(N) <= a constant (for all sufficiently small n)
2. GeneralCase: 
    T(N) <= aT(N/b) + O(N^d)
    where:
        a = number of recursive calls
        b = factor by which N shrinks on each reccursive call (> 1)
        d = amount of work done outside of reccursive calls (<=0)
```

```sorta js
T(N) = () => {
    if (a == b^d) {
        O(N^d*log(n))
    } else if (a < b^d) {
        O(N^d)
    } else if (a > b^d) {
        O(N^logb(a))
    }
}
```

## Examples

> Merge Sort:

```math
Number of recursive calls = 2
Factor by which N shrinks = 2
Amount of work done outside the recursive calls = 1 (linear)
```

2 == 2^1 so we match case 1 of the master method

```math
O(Nlog(N))

because d == 1
```

> Binary search

a = 1 (only recurse on half of the array)
b = 2 (half the array)
c = 0 (constant time with one comparison)

1 == 2^0 so case 1 `O(log(N))` because d was 0

## Some more explanation

a is the rate at which sub problems proliferate as you dive deeper into the recursion tree (evil)

b^d the rate at which we do less work at each level (good)
