package goodsadd

import (
	_ "cht/initial"
	"testing"
)

func NewGoodsAddRequest() *GoodsAddRequest {
	return &GoodsAddRequest{
		TotalNum:    1000,
		Name:        "wzl",
		Litpic:      "cht123",
		RedbagMoney: "1000",
		Content:     "4234",
	}
}

func TestAddGoods(t *testing.T) {
	gar := NewGoodsAddRequest()
	b := AddGoods(gar)
	if b == false {
		t.Fatalf("TestAddGoods failed")
	}
}
