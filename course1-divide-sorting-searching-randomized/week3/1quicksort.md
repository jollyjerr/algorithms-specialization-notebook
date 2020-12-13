# Quicksort

A slick algorithm with O(N logN) runtime

1. Pick a pivot element

2. Put elements less than the pivot to the left and greater than the pivot to the right

3. Two recursive calls to left and right bucket

Partitioning has O(N) time and _no extra memory partitioned_. Also cuts down the problem size.

Worst case runtime is O(N^2) if your pivot is the first element and the elements are already sorted. Bestl case runtime is O(N log(N)) if you luckily choose the median as every pivot.

The big idea: choose randomized pivots. A random pivot is "good enough often enough"

## Analysis

Can't apply the master method because of random, unbalanced subproblems

Remember worst case is O(N^2) and best case O(Nlog(N))

Key factor for analysis: number of comparisons that quicksort makes. Have to prove that the random number of comparisons is O(N log(N)) (`E[C]`)

We can do this with the "general decomposition principle"

### general decomposition principle

1. Identify random variable Y that you care about

2. Express Y as sum of indicator random variables `C = sigma(m, l = 1)*xl`

3. Apply linearity of expectation `E[Y] = sigma(m, l=1)*Pr[xl = 1]`

Notation:

```math
A = input array of n length

zi == ith smallest element of A
```

 Reduce to a family of smaller random variables:

 that identify the number of comparisons for a given pair of elements

```math
 For indices i < j (the two counters while quick sort is doing it's thing)
    xij(a) == number of times zi and zj are compared
```

xij is the random variable which evaluates to either 0 or 1

So by linearity of expectation:

```math
C(A) = sum(n-1)*sum(n)*xij(A)
```

> expectation of a sum equals the sum of expectations

### key claim

In quicksort - `Pr[zi,zj get compared] == 2/(j-i+1)`

3 possibilities are happening for a given set `zi, zi+1...zj-1, zj`

1. Something outside the set is chosen as a pivot, so they are not compared but passed to the same recursive call

2. Something inside the set is chosen, so they are not compared nor will they ever be compared

3. One of them is chosen, so they are compared

`2 (number of pivot choices that lead to comparison) / (j-i+1) (number of comparisons overall)`

### Result

`E(C) <= 2N * sum(n, k=2)1/k`

This second expression has an upper bound of `Intigral(1, n)1/x,d/x => log(x) => log(n) - log(1) => log(n)`

Expected number of comparisons results in `2N * log(n)`

suppress constants to result in `O(N log(N))`
