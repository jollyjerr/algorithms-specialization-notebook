# Sequence Alignment

Sum the penalties of all the flaws (gaps and mismatches) that result when you mutate two strings to have identical lengths and appear similarly

Pij = penalty of optimal alignment of Xi & Yj

```psyudo
A = 2D Array (letters of X and Y)
A[i,0], A[0,i] := i * cGap
for i =1 to m:
    for j=1 to n
        A[i,j] = min{
            A[i-1,j-1] + cxiyj,
            A[i-1,j] + cGap,
            A[i,j-1] + cGap,
        }
```

Runtime O(mn)

Traceback runtime O(m+n)
