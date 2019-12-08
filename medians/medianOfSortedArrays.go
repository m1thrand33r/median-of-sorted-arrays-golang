package medians

type MedianFinder struct {
	prev            int
	current         int
	positionTracker int
}

// TODO FIXME: assume the input is sorted, we could evaluate as we proceed through ach input
func (m *MedianFinder) FindMedianOfSorted(arrA []int, arrB []int) int {
	if len(arrA) > len(arrB) {
		return -1
	}

	n := len(arrA)
	var ptrA, ptrB = 0, 0
	m.positionTracker = 0
	m.prev, m.current = -1, -1

	// in our 2n element final array, the median elements are [n, n + 1], we need not traverse further than this
	medianRightEl := n + 1
	for m.positionTracker < medianRightEl {
		if ptrA == n {
			// process rest of arrB, we have consumed all of A according to the ptr
			m.processTail(medianRightEl, arrB, ptrB)
		} else if ptrB == n {
			// process rest of arrA, we have consumed all of B according to the ptr
			m.processTail(medianRightEl, arrA, ptrA)
		} else {
			// process the arrays, comparing the element of the sorted
			if arrA[ptrA] < arrB[ptrB] {
				m.setCurrentAndPrev(arrA[ptrA])
				ptrA++
			} else if arrB[ptrB] < arrA[ptrA] {
				m.setCurrentAndPrev(arrB[ptrB])
				ptrB++
			} else {
				// the elements are equal so we need to account for both
				m.setCurrentAndPrev(arrA[ptrA])
				ptrA++
				m.setCurrentAndPrev(arrB[ptrB])
				ptrB++
			}
		}
	}

	return (m.current + m.prev) / 2
}

func (m *MedianFinder) processTail(medianRightEl int, arrWTail []int, arrPtr int) {
	for m.positionTracker < medianRightEl {
		m.setCurrentAndPrev(arrWTail[arrPtr])
		arrPtr++
	}
}

func (m *MedianFinder) setCurrentAndPrev(newVal int) {
	// replace prev with current if required
	if m.current > 0 || m.current == -1 {
		m.prev = m.current
	}
	m.current = newVal

	// move the tracker forward
	m.positionTracker++
}
