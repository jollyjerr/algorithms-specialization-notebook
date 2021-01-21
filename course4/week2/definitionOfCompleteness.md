# Definition Of Completeness

A computational problem is "easy" if, given a polynomial time algorithm, you can reduce to the solution you want.

Contrapositive: if neither problems can be solved in polynomial time.

## Definition of NP

"nondeterministic-polynomial" but this is a bad definition

"possibly exponential time"

Roughly, all problems you could solve using brute force search

Meet the following criterial:

1. Solutions always have length polynomial in the input size

2. purported solutions can be verified in polynomial time

If you REALLY want to show something is NP complete - prove that if it was solved, you could use reduction to solve thousands of unsolved computational problems

### At the point of exasperation - ask if it is intractable

Settle on an np-complete problem, and prove that it can be reduced to your problem. QED.

## P != NP

Just because a problem is solvable and np does NOT mean it is in P???? -> can be solved by a polynomial time algorithm??? We are not sure

Traveling salesman problem is NP but not P solvable. Why? We have no proof :(

## How to approach NP complete problems

Align expectations: the solution will be hard and you will need to make compromises

### Strat 1

1. Focus on computationally tractable special cases within your domain

2. Leverage domain knowledge to crucial elements

3. Mix brute force search with computationally tractable special cases

### Strat 2

Move to heuristics

1. Accept algorithms that are sometimes wrong

### Strat 3

Solve in exponential time but smarter than brute force search
