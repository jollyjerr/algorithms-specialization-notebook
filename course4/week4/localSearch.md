# Local Search

Let X = set of candidate solutions to a problem.

key ingredient: neighborhoods.

- for each x in X, specify which y in X are its "neighbors"

x,y are neighbors if they differ by one element

## Once you settled on a neighborhood

Pick a solution in the neighborhood, [will discuss later]

look for a neighbor solution y that is better, move to y if it is found

return when you cannot iterate anymore

## Decisions you need to make

1. How to pick initial solution x?

If you have 0 idea how to even get close -> run local search millions of times with random decisions.

If you have a good idea of a close to optimal solution -> use a local-search post-processing step. This can only make your solution better.

2. If you have multiple superior neighboring y, which to choose?

No real answer -> this is a second opportunity to inject randomness and rerun the algorithm millions of times.

Work hard with domain knowledge.

3. What are the neighborhoods? How big are they?

The bigger the neighborhood, the more time you have to invest in the algorithm, but fewer local optima


