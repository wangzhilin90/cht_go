package goodsedit

import (
	_ "cht/initial"
	"testing"
)

func NewGoodsEditRequestStruct(id int32, money string) *GoodsEditRequestStruct {
	return &GoodsEditRequestStruct{
		ID:          id,
		RedbagMoney: money,
	}
}

func TestEditGoods(t *testing.T) {
	ger := NewGoodsEditRequestStruct(12, "60.01")
	ges := goodseditservice{}
	res, _ := ges.EditGoods(ger)
	if res.Status != EDIT_GOODS_SUCCESS {
		t.Fatalf("TestEditGoods failed")
	}
}
