package goodsedit

import (
	_ "cht/initial"
	"testing"
)

func NewGoodsEditRequest(id int32, money string) *GoodsEditRequest {
	return &GoodsEditRequest{
		ID:          id,
		RedbagMoney: money,
	}
}

func TestEditGoods(t *testing.T) {
	ger := NewGoodsEditRequest(12, "50.00")
	b := EditGoods(ger)
	if b == false {
		t.Fatalf("TestEditGoods failed")
	}
}
