package coins

import (
	"testing"
)

var total int = 33
var values []int = []int{2,5,10}

func TestGetMinCoinCount01(t *testing.T) {
	t.Log(getMinCoinCount01(total, values))
}

func TestGetMinCoinCount02(t *testing.T) {
	t.Log(getMinCoinCount02(total, values))
}

func TestGetMinCoinCount03(t *testing.T) {
	t.Log(getMinCoinCount03(total, values))
}

func TestGetMinCoinCount04(t *testing.T) {
	t.Log(getMinCoinCount04(total, values))
}
