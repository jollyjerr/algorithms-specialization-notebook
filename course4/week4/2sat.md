# The 2-SAT Problem

Input:

1. n Boolean variables `x1,x2,...,xn`

2. m clauses of 2 literals each.

Output:

Is it or is it not possible to set boolean values so that every m clause is satisfied simultaneously.

## Info

2-SAT: Can be solved in polynomial time!

3-SAT: canonical NP-Complete problem

## Papadimitriou's 2-SAT Algorithm

```psyudo
Repeat log2 n times:
    choose random initial assignment
    repeat 2n^2 times:
        if current assignment satisfies all clauses, halt and report;
        else, pick arbitrary unsatisfied clause and flip the value of one of its variables [choose between the two uniformly at random]
report "unsatisfiable";
```
