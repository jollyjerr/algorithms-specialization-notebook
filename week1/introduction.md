# Introduction

The textbook resource for the class is [Algorithms Illuminated (Part 1) (by Tim Roughgarden)](https://www.amazon.com/dp/0999282905)

If struggling with math concepts, see: [Mathematics for Computer Science (by Eric Lehman and Tom Leighton)](https://www.cs.princeton.edu/courses/archive/fall06/cos341/handouts/mathcs.pdf)

## Integer Multiplication

Many algorithms can be outlined by looking at an *input* and an *output* first, and then defining the process the completes that transformation. You can then asses the *performance* of the algorithm by looking at the number of basic operations that it performs as a function of the length of the input numbers.

```pseudo
Input: two n-digit numbers x and y
Output: the product of x * y
```
A *correct* algorithm will always produce the right answer regardless of input.

>They say the following, perhaps the most important principle of all, for the good algorithm designer is to refuse to be content. - Hopcraft and Aho

Always ask "can we do better?"
