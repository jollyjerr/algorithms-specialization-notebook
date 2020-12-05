# Asymptotic Analysis

"Big-Oh" means that you have reduced F(N) to the sweet spot of abstraction by performing Asymptotic Analysis. *Suppress* constant factors (from the environment EX: lines of code) and lower-order terms (that are irrelevant for large inputs).

From the merge sort example that we proved (via binary recursion tree) to have a upper bound of 6N(log2(n) + 6N), you can suppress constant factors and lower-order terms even farther to reach a "Big-Oh" of `Nlog(N)`

> asymptotic analysis is for comparing two fundamentally different ways of salving a problem

## Big O

A mathematical function can be Big O of another function if it is bounded above (less than) by that function after a certain point and multiplied by a constant C (AKA less than or equal to)

Just have to prove that one constant C works!

```
T(N) is O(F(N)) if and only if C * F(N) is forevermore above T(N) after N * x (the crossing point)
```

Close relatives are Omega (greater than or equal to) and Theta (equal to)

```
T(N) is BigOmega(F(N)) if and only if C * F(N) is forevermore below T(N) after N * x
```

Theta means both are true

Little o means all constant C there exists some larger "N not" there is a upper bound C *f(n) for T(N) `strictly less than`
