# Counting Inversions

Inversions: for an array of N entries and in inversion is when `i < j && array[i] > array[j]`

Useful for collaborative filtering. The more inversions between two arrays, the more dissimilar the two arrays are - so more similar/adjacent arrays show similar preferences.

The largest number of inversions an array can have is `(n(n-1))/2`

Simplest solution - Brute Force: O(n^2)

```pseudo
for array i:
    for array j = i+1:
        if array[i] > array[j]:
            counter++
```

For a divide and conquer approach

1. Split the array

2. Recursively compute left and right inversions

3. Merge result and account for "split" inversions (right in the middle of the array and will be missed by left and right counts)

There can be as many as N^2 split inversions

Claim to solve split inversion sub routine in non quadratic time:

> The split inversions that involve an element of the second half of the array are precisely those elements remaining in the first array when that element gets copied over to the output array.

Notes:

If you add O(N) to O(N) N time is is not O(N), but if you add O(N) to O(N) a constant amount of times it *is* O(N)
