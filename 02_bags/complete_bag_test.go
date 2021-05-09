package bags

import (
"testing"
)

func TestCompleteBag(t *testing.T) {
	w := []int{3, 2, 1}
	v := []int{5, 2, 3}
	N := len(w)
	W := 5
	t.Log(completeBag01(W, N, w, v))
	t.Log(completeBag02(W, N, w, v))
	t.Log(completeBag03(W, N, w, v))
}

//func TestLastStoneWeightII(t *testing.T) {
//	stone := []int{2,7,4,1,8,1}
//	t.Log(lastStoneWeightII(stone))
//}


