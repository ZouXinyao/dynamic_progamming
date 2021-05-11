package subsequence

import "testing"

func TestLongestPalindromeSubseq(t *testing.T) {
	s := "cbbd"
	t.Log(longestPalindromeSubseq(s))
}

func TestLongestCommonSubsequence(t *testing.T) {
	s1 := "bsbininm"
	s2 := "jmjkbkjkv"
	t.Log(longestCommonSubsequence(s1, s2))
}

