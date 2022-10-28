# Coding Assessment

## Problem 2

Amazon Web Services (AWS) has several data centers which have multiple processors that perform computations. In one such data center, these processors are placed in a sequence *with their IDs denoted by 1,2,…,n*. Each processor consumes a certain amount of power to boot up, denoted by `bootingPower[i]`. After booting, a processor uses `processingPower[i]` of power to run the processes. For maximum utilization, the data center wishes to group these processors into clusters. Clusters can only be formed of processors located adjacent to each other. For example, processors 2, 3, 4, 5 can form a cluster, but 1, 3, 4 cannot.

The net power consumption of a cluster of `k` processors (`i`,`i+1`,…,`i+k−1`) is defined as:

<p align="center">
  <img src="../assets/cq2-1.png">
</p>

That is, *net power consumption* = *maximum booting power* among the `k` processors + (*sum of processing power of processors*) * `k`.

A cluster of processors is said to be *sustainable* if its *net power consumption* does not exceed a given threshold value `powerMax`.

Given the booting power consumption and the processing power consumption of `n` processors denoted by `bootingPower` and `processingPower` respectively, and the threshold value `powerMax`, find the maximum possible number of processors which can be grouped together to form a *sustainable* cluster. If no such clusters can be formed, return 0.

Here,
* Maximum booting power = max(3, 6, 1) = 6
* Sum of processing powers = 2 + 1 + 3 = 6
* Thus, net power consumption = 6 + 6 * 3 = 24 ≤ powerMax

Thus, we can group `k=3` processors to form a *sustainable* cluster. Note that the minimum power consumption to form a cluster of `k=4` processors is 46, by forming a cluster of the first 4 processors. Since this cost is greater than the threshold, we cannot form a cluster with 4 processors. The maximum possible cluster size is 3.

**Function Description**

Complete the function `findMaximumSustainableClusterSize` in the editor below.

`findMaximumSustainableClusterSize` has the following parameters:
* `int processingPower[n]`: the power consumption of each processor in running the processes
* `int bootingPower[n]`: the power consumption of each processor in bootup
* `long_int powerMax`: the threshold power consumption

*Returns*
* `int`: the maximum cluster size which is *sustainable*, if no cluster exists, return 0.

**Constraints**
* `1 ≤ n ≤ 10^5`
* `1 ≤ powerMax ≤ 10^14`

**Note**: It is not mandatory for **all** clusters of size `k` to be *sustainable*. Even one such cluster is sufficient.

**Example**

```
processingPower = [2, 1, 3, 4, 5]
bootingPower = [3, 6, 1, 3, 4]
powerMax = 25
```

If `k=2`, any adjacent pair can be chosen. The highest usage is the pair `[4,5]` with net power consumption 4+(4+5)*2=22. Next, try `k=3`. Group the first 3 processors together as:

```
Maximum Power = 25
Cluster of 3 processors
Maximum booting power = 6
Sum of processing power = 6
Net Power Consumption = 6 + 6 * 3 = 24
```

<p align="center">
  <img src="../assets/cq2-2.png">
</p>

## Solution

This problem can be approached as a sliding window problem given the constraint:

>Clusters can only be formed of processors located adjacent to each other.

Therefore, starting at index zero, the window is adjusted such that the maximum cluster is formed when the following criteria is met:

```
net power consumption = maximum booting power + (sum of processing power) * k
```

where net power consumption ≤ max power (`powerMax`).

This approach is suboptimal, since it uses a double nested loop. That is, the time complexity is `O(n^2)`.

Instead, this problem can be solved in linear time (`O(n)`) using a monotonic queue (or monotone priority queue), a variant of a priority queue in which the priorities of extracted items are required to form a monotonic sequence (i.e. for all `n ∈ N` a<sub>n+1</sub> ≥ a<sub>n</sub> for monotonically increasing and a<sub>n+1</sub> ≤ a<sub>n</sub> for monotonically decreasing).

The monotonic queue is created using a deque (pronounced *deck*), for which elements can be added to or removed from either the front (head) or back (tail) in constant (`O(1)`) time.

The application of the monotonic queue is used to maintain the maximum booting power of the current group of processors. If we did not use this data structure, the maximum booting power would need to be calculated for each window, thus necessitating the double nested loop of the naive solution.
