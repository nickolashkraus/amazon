package pkg

/*
 * Complete the 'findMaximumSustainableClusterSize' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY processingPower
 *  2. INTEGER_ARRAY bootingPower
 *  3. LONG_INTEGER powerMax
 */

import (
	"github.com/gammazero/deque"
)

func findMaximumSustainableClusterSizeNaive(processingPower []int32, bootingPower []int32, powerMax int64) int32 {
	// This problem can be approached as a sliding window problem given the
	// constraint:
	//
	//   > Clusters can only be formed of processors located adjacent to each
	//     other.
	//
	// Therefore, starting at index zero, the window is adjusted such that the
	// maximum cluster is formed when the following criteria is met:
	//
	//   net power consumption = maximum booting power + (sum of processing power) * k
	//
	// where net power consumption ≤ max power.
	maxClusterSize := 0
	numProcessors := len(processingPower)
	for i := 0; i < numProcessors; i++ {
		clusterSize := 1
		sumProcessingPower := processingPower[i]
		// Assume first element has max booting power.
		maxBootingPower := bootingPower[i]
		for j := i + 1; j < numProcessors; j++ {
			if bootingPower[j] > maxBootingPower {
				maxBootingPower = bootingPower[j]
			}
			sumProcessingPower += processingPower[j]
			power := maxBootingPower + sumProcessingPower*int32(clusterSize)
			if int64(power) <= powerMax {
				clusterSize++
			}
			if clusterSize > maxClusterSize {
				maxClusterSize = clusterSize
			}
		}
	}
	return int32(maxClusterSize)
}

func findMaximumSustainableClusterSizeOptimal(processingPower []int32, bootingPower []int32, powerMax int64) int32 {
	// A monotonic queue (or monotone priority queue) is maintained with the
	// maximum booting power of the current group of processors.
	//
	// The monotonic queue is implemented using a deque, double-ended queue, in
	// order to efficiently add and remove items at either end with O(1)
	// performance.
	//
	// Pointers to the start and end of the sliding window are maintained in
	// order to preserve the following criteria:
	//
	//   net power consumption = maximum booting power + (sum of processing power) * k
	//
	// where net power consumption ≤ max power.
	//
	// NOTE: In order to preserve the bounds of the sliding window, the deque
	// holds the indices of bootingPower, not its values.
	mq := deque.New[int]()
	maxClusterSize, clusterSize, start, end := 0, 0, 0, 0
	var sumProcessingPower int32 = 0
	for end < len(processingPower) {
		// Remove elements from the back of the queue until the queue is empty or
		// the element at the back of the queue is greater than the current
		// element, then append the current element to the back of the queue.
		for mq.Len() > 0 && bootingPower[mq.Back()] < bootingPower[end] {
			mq.PopBack()
		}
		mq.PushBack(end)
		maxBootingPower := bootingPower[mq.Front()]
		sumProcessingPower += processingPower[end]
		clusterSize++
		power := maxBootingPower + sumProcessingPower*int32(clusterSize)
		if int64(power) <= powerMax {
			end++
		} else {
			sumProcessingPower -= processingPower[start]
			clusterSize--
			start++
			end++
		}
		if clusterSize > maxClusterSize {
			maxClusterSize = clusterSize
		}
		// Preserve the bounds of the sliding window by checking that the element
		// at the front of the queue is still within the window. If not, remove it.
		if start > mq.Front() {
			mq.PopFront()
		}
	}
	return int32(maxClusterSize)
}
