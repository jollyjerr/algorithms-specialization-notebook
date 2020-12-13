# Probability

## Sample space

big omega notation, "all possible outcomes"

Rolling two dice:

```math
{(1,1),(1,2)(2,2)...(6,6)}
```

## Events

A subset of sample space. Add the probability of each individual things.

## Random variables

A real-valued function defined on the sample space omega.

Ex: a running time of a random based algorithm || size of sub array passed into 1st recursive call

### Expectation notation

E[x] of x = average value of x

## Linearity of expectation

> the expected value of a sum of random variables is equal to the sum of the expectations of the individual random variables

```math
E[sum(xi)] === sum(E[xi])
```

Holds true even if the random variables are not independent

Use it to find E[x] for rolling two dice

```math
1 dice = 1/6 || 3.5

3.5 + 3.5 => 7
```

## Conditional probability

X, Y represent two events, then Pr[X|Y] ("x given y") === `Pr[intersectionOf(X, Y)]/Pr[Y]`

## Independence of events

X and Y are independent if and only if `Pr[intersectionOf(x, y)] === Pr[x] * Pr[y]`

Knowing that one happened gives you no new information on if the other has happened.
