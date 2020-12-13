# Randomized Selection

Reduction - when you solve one problem by reducing it to another problem that you know how to solve.

You can reduce randomized selection to an `O(n log(n))` solution by

1. sorting the input array

2. The `ith` statistic will then be the `ith` element in the array

## Analysis

RSelect uses <= Cn operations outside of the recursive call (for some constant c > 0)

Notation: RSelect is in `phasej` if current array size is between `((3/4)^j+1) * n` and `((3/4)^j) * n`

`j` quantifies the number of times we have made 75% progress.

`Xj` = number of recursive calls during phase j

Note: The running time of RSelect `<= sum(phasesj)*Xj*c*(((3/4)^j) * n)`. These values all come from above and make a clean upper bound

Both sides of that equation are random values so you do calculations with average values.

50% chance of picking a pivot that gives you at least a 25/75 split - which puts you in phase j

`E[Xj] <= expected number of fair coin flips to get one heads` (geometric random variable (1/2))

Coin flip can be represented as `E[N] = 1 + 1/2 * E[N]` which forces `E[N]` to be 2

so... `E[Xj] <= 2`

so you can rework the problem using linearity of expectation 

```math
rSelect <= Cn * sum(phasej) * ((3/4)^j) * E[Xj]

2cn * sum(phasej) * ((3/4)^2)

use geometric sum

<= 8cn
```
