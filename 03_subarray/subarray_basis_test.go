package subarray

import "testing"

func TestCountSubstrings(t *testing.T) {
	s := "aabcabcba"
	t.Log(countSubstrings(s))
}

func TestLongestPalindrome(t *testing.T) {
	s := "aabcabcba"
	t.Log(longestPalindrome(s))
}

func TestMaxSubArray(t *testing.T) {
	arr := []int{-2,1,-3,4,-1,2,1,-5,4}
	t.Log(maxSubArray(arr))
}