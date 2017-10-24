package goodslist

import (
	_ "cht/initial"
	"testing"
)

func NewGoodsListRequestStruct(limitNum, LimitOffset int32) *GoodsListRequestStruct {
	return &GoodsListRequestStruct{
		LimitNum:    limitNum,
		LimitOffset: LimitOffset,
	}
}

func TestGetGoodsList(t *testing.T) {
	glrs := NewGoodsListRequestStruct(2, 1)
	gls := goodslistservice{}
	res, _ := gls.GetGoodsList(glrs)
	if res.Status != GET_GOODS_LIST_SUCCESS {
		t.Fatalf("TestGetGoodsList query failed")
	}
	t.Logf("TestGetGoodsList return value:%v", res)
}
