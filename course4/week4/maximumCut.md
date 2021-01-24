# The Maximum Cut Problem

> For use in understanding the Local Search approach to NP-Complete problems

Input: Undirected graph G

Goal: a cut (A,B) - a partition of V into two non-empty sets - that maximizes the number of crossing edges.

Sad Fact: NP-complete.

## Local Search

Small improvements to a guess

1. Let (A,B) be an arbitrary cut of G

2. While there is a vertex v with `dv(A,B) > Cv(A,B)`: move v to other side of the cut

[increases number of crossing edges]
