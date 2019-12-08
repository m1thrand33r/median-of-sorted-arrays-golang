package medians

import "testing"

func TestFindMedianOfSorted(t *testing.T) {
	arrA := []int{1, 12, 15, 26, 38}
	arrB := []int{2, 13, 17, 30, 45}

	finder := MedianFinder{}

	medianOfSorted := finder.FindMedianOfSorted(arrA, arrB)
	if medianOfSorted != 16 {
		t.Errorf("We expected %d but got %d", 16, medianOfSorted)
	}
}
