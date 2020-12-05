# Merge Sort

Good introduction to divide and conquer:

- Improves over Selection, Insertion, Bubble sort which all have n^2 or quadratic running time

Divide and Conquer algorithms:

1. Divide a hard problem into many small problems that can be solved recursively
2. Solve them
3. Somehow merge the results

Merge sort is the most transparent application of this tactic

A note on guiding principles for algorithmic analysis:

- worst-case: look for guarantees of performance on running time that hold for every possible input on a given size

- Asymptotic: remain focused on the rate of growth of an algorithms performance rather than low-order terms or small changes to constant factors

Merge sort will be analyzed by a <mark>recursion tree</mark> method which generalizes greatly for divide and conquer algorithms

Note: merge sort in this lecture is for lists of distinct numbers

## Analysis

merge sub routine: For each N entries you are doing at most 4N + 2 (for even number N lists) or 6N if you are being sloppy

Balance of *explosion of sub problems* and each subproblem has a *smaller input*

**Claim:** Merge sort needs at most a constant (in this case 6) times N times log2(n) which is better than the quadratic n^2 of bubble sort etc...

Log2(n) -> <mark>A curve that grows very flat, very quickly, as N grows large</mark>

### Proof of claim (via a binary recursion tree method)

```pseudo
            level0 (root): [6,5,4,3,2,1]
                                |
            level1      [6,5,4] - [3,2,1]
etc....... log2(n)
```

Bottoms out in base-cases when arrays are only of size 0 or 1

**The total number of levels will always be log2(n) + 1**

At a given level j on the tree there are 2j problems each with N/2j arguments

Compute the amount of work at each level j -> ends up being an upper bound 6N independent of level

So total upward bound comes out to be 6N(log2(n) + 6N)

Thus proving this is more efficient (in terms of time complexity) to an iterative solution
