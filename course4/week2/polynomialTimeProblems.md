# NP-Completeness - P: Polynomial-Time Solvable Problems

Sad fact: many important problems seem impossible to solve efficiently

We formalize computational intractability using "NP-Completeness"

## What is computational tractability?

A problem is polynomial-time solvable if there is an algorithm that correctly solves it in O(n^k) time, for some constant k.

Some mathematical evidence suggests that randomized algorithms can not take an unsolvable problem and make it solvable

## The class P

P = the set of poly-time solvable problems

Membership in P is a test for computation tractability

> Problems in P can be solved with off the shelf solutions, problems outside of P typically require lots of resources and domain knowledge!

Example: Traveling salesman problem

Visit every node in a connected undirected graph once and minimize the edge costs. This has been studied for 60+ years with no known polynomial time solution.
