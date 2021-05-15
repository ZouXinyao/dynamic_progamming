package subsequence

import "testing"

func TestMaxLongestRiseSubque(t *testing.T) {
	s := []int{3, 3, 56, 2, 5, 8, 10, 2}
	t.Log(maxLongestRiseSubque(s))
}
