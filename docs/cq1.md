# Coding Assessment

## Problem 1

Imagine you are shopping on Amazon.com for some good weight lifting equipment. The equipment you want has plates of many different weights that you can combine to lift.

The listing on Amazon gives you an array, `plates`, that consists of `n` different weighted plates, in kilograms. There are no two plates with the same weight. The element `plates[i]` denotes the weight of the `i`<sup>th</sup> plate from the top of the stack. You consider weight lifting equipment to be good if the plate at the top is the lightest, and the plate at the bottom is the heaviest.

More formally, the equipment with array `plates` will be called good weight lifting equipment if it satisfies the following conditions (assuming the index of the array starts from 1):
* `plates[1] < plates[i]` for all (`2 ≤ i ≤ n`)
* `plates[i] < plates[n]` for all (`1 ≤ i ≤ n−1`)

In one move, you can swap the order of adjacent plates. Find out the minimum number of moves required to form good weight lifting equipment.

**Example**

Let the plates be in the order:

```
plates = [3, 2, 1]
```

In the first move, we swap the first and the second plates. After swapping, the order becomes:

```
plates = [2, 3, 1]
```

In the second move, we swap the second and the third plates. After swapping, the order becomes:

```
plates = [2, 1, 3]
```

In the third move, we swap the first and the second plates. After swapping, the order becomes:

```
plates = [1, 2, 3]
```

Now, the array satisfies the condition after 3 moves.

**Function Description**

Complete the function `getMinMoves` in the editor below.

`getMinMoves` has the following parameter:
* `int plates[n]`: the distinct weights

*Returns*
* `int`: the minimum number of operations required

**Constraints**
* `2 ≤ n ≤ 10^5`
* `1 ≤ plates[i] ≤ 10^9` for all (`1 ≤ i ≤ n`)
* `plates` consists of distinct integers.
