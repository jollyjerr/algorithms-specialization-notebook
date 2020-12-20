# Data Structures continued

## Hash table

Basically an array with it's positions indexed by the keys you are storing

### Purpose

Maintain a possibly evolving set of stuff.

### Operations

Insert

Delete

Lookup

all in constant time

### Implementation

1. Identify universe of application - (what type of data, omega set, stuff like that)

2. Choose a hash function -> responsible for translating things we care about into positions in our array. Takes in a key and spits out position

3. Store each item where the hash function tells you to

4. Deal with collisions

- Separate chaining: Revert to list based solution on collisions. Each bucket contains a list

- Open addressing: Replace has function with hash "probe sequence". Retry hashing again and again when you hit collisions

A hash tables "load factor" is #ofobjects/#ofbuckets

#### What makes a good hash function?

It leads to good performance - constant time and spreads the data out

This is a hard problem, here is a quick and dirty solution:

```math
object -> integer -> bucket

First arrow is "hash code transformation"
Second arrow is "compression"
```

A perfect/universal hash function does not exist

For every hashing function, you can develop a pathological data set that will reduce the performance to full O(N) scans

You can combat that with cryptographic hash functions or designing a family of hash functions and picking a random one at runtime

## Bloom filters

Just store if we have seen something. Very space efficient but with small false positive probability and no deletions.

Hash functions at number `k` to point object to array of n bits

Just set bits to 1 where hash function points you too
