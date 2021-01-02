# Huffman Codes

> Find the optimal binary prefix-free encoding of a set of characters (Optimal compression)

Binary codes are interchangeable with binary trees (0 left, 1 right)

Note: Prefix free trees will only have symbols on the leaves

## To decode

repeatedly follow path from root (following instruction in message) until you hit a leaf

The steps needed to decode/encode a symbol is the depth of a given leaf

## Problem definition

```math
Given T = tree with leaves == symbols of the set SIGMA

then L(T) // average encoding length == depth of i in T for i in T * frequency of i in T
```

### Building a tree

Start with the leaves of the tree, and build the tree from bottom-up using successive merges

When can we be sure that a merge is safe??

1. The symbols that are the least frequent should merge first. (because merging increases the cost)

2. The merges symbols now have a summed frequency of p1 + p2.

3. Run the algorithm again with this new meta-symbol as a single entry

4. After base case, split all of the merged nodes and assign 0 to left and 1 to right

```psyudo
if len(T) == 2 return [0, 1]
else take 2 symbols with smallest frequencies A and B and merge to ab
recurse
Split leafs
return
```

If you use heaps, reinsert the meta node after popping off the top

You could also use two queues and sorting upfront
