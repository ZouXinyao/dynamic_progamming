package bags

import (
"testing"
)

func TestZeroOneBag(t *testing.T) {
	w := []int{3, 2, 1}
	v := []int{5, 2, 3}
	N := len(w)
	W := 5
	t.Log(zeroOneBag(W, N, w, v))
}

func TestLastStoneWeightII(t *testing.T) {
	stone := []int{2,7,4,1,8,1}
	t.Log(lastStoneWeightII(stone))
}


