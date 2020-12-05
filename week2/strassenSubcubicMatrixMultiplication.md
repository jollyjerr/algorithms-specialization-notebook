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


