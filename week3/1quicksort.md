# Quicksort

A slick algorithm with O(N logN) runtime

1. Pick a pivot element

2. Put elements less than the pivot to the left and greater than the pivot to the right

3. Two recursive calls to left and right bucket

Partitioning has O(N) time and _no extra memory partitioned_. Also cuts down the problem size.

Worst case runtime is O(N^2) if your pivot is the first element and the elements are already sorted. Bestl case runtime is O(N log(N)) if you luckily choose the median as every pivot.

The big idea: choose randomized pivots. A random pivot is "good enough often enough"
