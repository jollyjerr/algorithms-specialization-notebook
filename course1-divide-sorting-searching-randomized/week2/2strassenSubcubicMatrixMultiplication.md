# Strassen's Subcubic Matrix Multiplication

The problem:

Two matrixes (X and Y), all with the same dimensions N X N

Multiply those two to result in a third matrix (Z) where:

```math
Zij == (ith row of X) (dot product) (jth column of Y)
```

We are using N to denote the dimensions, so there are N^2 arguments. The best we could possibly hope for is a running time of O(N^2)

 
Simple example where N == 2

```math
(a b)(e f) === (ae+bg af+bh)
(c d)(g h)     (ce+dg cf+dh)
```

An iterative solution will run in cubic (n^3) time.

## Divide and conquer solution

### Try #1

Break the matrixes into 4 square blocks and recursively solve (notice that you cannot simply break the matrixes in half because they will no longer be square)

```math
(a b)(e f) === (ae+bg af+bh)
(c d)(g h)     (ce+dg cf+dh)
```

Turns out that this solution is cubic (N^3) - exactly the same as the iterative approach

### Try #2

Strassen's algorithm comes in to reduce the number of recursive calls (from try #1) from 8 down to 7.

Break the two matrixes into 7 cleverly chosen products:

```math
x = (a b)      y = (e f)
    (c d)          (g h)

P1 = a(f-h)
P2 = (a+b)h
P3 = (c+d)e
P4 = d(g-e)
P5 = (a+d)(e+h)
P6 = (b-d)(g+h)
P7 = (a-c)(e+f)

Compute X * Y as such:

(P5+P4-P2+P6  P1+P2)
(P3+P4        P1+P5-P3-P7) // The long expressions work out with cancellations
```
