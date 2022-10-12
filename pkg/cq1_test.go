package pkg

import (
	"testing"
)

func TestGetMinMovesNaive(t *testing.T) {
	plates := []int32{3, 2, 1}
	ret := getMinMovesNaive(plates)
	if ret != 3 {
		t.Fail()
	}
	plates = []int32{1, 20, 6, 4, 5}
	ret = getMinMovesNaive(plates)
	if ret != 5 {
		t.Fail()
	}
	plates = []int32{10, 3, 4, 2, 5, 7, 9, 11}
	ret = getMinMovesNaive(plates)
	if ret != 8 {
		t.Fail()
	}
}

func TestGetMinMovesOptimal(t *testing.T) {
	plates := []int32{3, 2, 1}
	ret := getMinMovesOptimal(plates)
	if ret != 3 {
		t.Fail()
	}
	plates = []int32{1, 20, 6, 4, 5}
	ret = getMinMovesOptimal(plates)
	if ret != 5 {
		t.Fail()
	}
	plates = []int32{10, 3, 4, 2, 5, 7, 9, 11}
	ret = getMinMovesOptimal(plates)
	if ret != 8 {
		t.Fail()
	}
}
