package pkg

import (
	"testing"
)

func TestFindMaximumSustainableClusterSizeNaive(t *testing.T) {
	processingPower := []int32{2, 1, 3, 4, 5}
	bootingPower := []int32{3, 6, 1, 3, 4}
	powerMax := int64(25)
	ret := findMaximumSustainableClusterSizeNaive(processingPower, bootingPower, powerMax)
	if ret != 3 {
		t.Fail()
	}
}

func TestFindMaximumSustainableClusterSizeOptimal(t *testing.T) {
	processingPower := []int32{2, 1, 3, 4, 5}
	bootingPower := []int32{3, 6, 1, 3, 4}
	powerMax := int64(25)
	ret := findMaximumSustainableClusterSizeOptimal(processingPower, bootingPower, powerMax)
	if ret != 3 {
		t.Fail()
	}
}
