# Dynamic Programming

> Learning by developing an algorithm - takes practice to perfect

## The problem

Input: A Path graph `G = (V, E)` with nonnegative weights on vertices

Output: Subset of nonadjacent vertices - an independent set - of maximum total weight

"The max weight independent set problem"

### How can we approach this?

1. Brute force search - iterate over all combinations and remember the max one. obviously correct but exponential runtime

2. Greedy approach - Each step pick the vertex with the highest weight and rule out adjacency. Not always correct :(

3. Divide and Conquer - Break the path into two paths, recursively compute a max independent set in each, combine results. Difficult to implement and quadratic time.

### A new paradigm

Systematic Critical Step: reason about the structure of an optimal solution in terms of optimal solutions of smaller subproblems

The optimal solution either includes the last vertex or does not. Try both possibilities and see which one is better? Sounds like a weird  version of brute force search...

### Turn it into a linear time algorithm

There are only O(N) distinct sub problems

Every step is guaranteed to be handed a prefix of the original graph

Simply memoize the answer to each subproblem

```psyudo
Let Gi = 1st i vertices of G
A[0] = 0, A[1] = w1
for i -> n:
    A[i] = maxOf{A[i-1], A[i-2] + wi}
```

#### Reconstruct the optimal solution

Trace back through filled-in array to reconstruct optimal solution

Check if case 1 or 2 won

```psyudo
let S = null
while i >= 1
    if A[i-1] > A[i-2]:
        i--
    else:
        add vi to S, i -= 2
```

## Principles of Dynamic Programming

1. Identify a suitable collection of subproblems - the lower bound of your algorithm

2. Make it so "larger" subproblems are made up of solutions of "smaller" subproblems

3. Be able to answer the original question

Best case is if the answer to a larger subproblem can be expressed as a function of the solution of smaller subproblems.

Dynamic programming came out of 1950's and was named that way as a fancy version of "Math planning"
