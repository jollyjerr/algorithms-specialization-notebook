# Greedy Algorithms

There is no magic potion to cure computational problems - you have to learn different paradigms.

Some big ones

1. Divide and conquer

2. Randomized algorithms

3. Greedy algorithms

4. Dynamic programming

## What is a greedy algorithm?

Makes a sequence of myopic decisions and hope everything works out at the end.

Example: Dijkstra

Typically have easy running time analysis

Typically have difficult correctness proofs or are not even correct

### Proof of correctness

Proving correctness of greedy algorithms is an art

1. Induction on the decisions made by the algorithm

2. "Exchange argument" - exchange an optimal solution with your greedy solution ("farthest-in-the-future" impossible algorithm being used as a benchmark agains real optimal caching algorithms)

3. Whatever works!

## Greedy algorithm for scheduling jobs

A jobs `completion time` is the length of itself plus all other jobs that came before it

An `objective function` decides what to prioritize. One popular one is to minimize the `sum of completion times * their weight`

Note: Order by decreasing ratio `weight/length` (decrease to ratios and then sort by ratio)

## Building a greedy solution

1. Focus on some special cases of the problem where it is clear how you should proceed

2. Resolve conflicting intuitions about the special cases

3. Rule out proposed algorithms by breaking them (as simple as possible but no simpler)
