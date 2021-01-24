# A Greedy Knapsack Heuristic

In the best case scenario - when you provide a heuristic you should provide a performance guarantee.

For knapsack:

Sort and reindex so items are: v/w >= v/w.....n

Read and take what you can and leave what does not fit

> This algorithm can be arbitrarily bad

You can improve it by checking your answer against each item and if one item overpowers the answer found in the algorithm. This fix turns the algorithm to a 1/2 approximation algorithm. The answer will always be at least 50% the actual answer.

Runs in `O(Nlogn)`

One way to improve this heuristic is to limit item sizes. If each item is w <= 10% of the knapsack then you will hit 90% accuracy. <= 1% and you hit 99%
