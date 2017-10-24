package goodsdel

import (
	_ "cht/initial"
	"testing"
)

func NewGoodsDeLRequestStruct(id int32) *GoodsDeLRequestStruct {
	return &GoodsDeLRequestStruct{
		ID: id,
	}
}

func TestDelGoods(t *testing.T) {
	gdrs := NewGoodsDeLRequestStruct(117)
	gds := goodsdelservice{}
	res, _ := gds.DelGoods(gdrs)
	if res.Status != DELETE_GOODS_SUCCESS {
		t.Fatalf("TestDelGoods failed")
	}
}
